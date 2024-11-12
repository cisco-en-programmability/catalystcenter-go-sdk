package catalyst

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type PathTraceService service

type RetrievesAllPreviousPathtracesSummaryV1QueryParams struct {
	PeriodicRefresh bool    `url:"periodicRefresh,omitempty"` //Is analysis periodically refreshed?
	SourceIP        string  `url:"sourceIP,omitempty"`        //Source IP address
	DestIP          string  `url:"destIP,omitempty"`          //Destination IP address
	SourcePort      float64 `url:"sourcePort,omitempty"`      //Source port
	DestPort        float64 `url:"destPort,omitempty"`        //Destination port
	GtCreateTime    float64 `url:"gtCreateTime,omitempty"`    //Analyses requested after this time
	LtCreateTime    float64 `url:"ltCreateTime,omitempty"`    //Analyses requested before this time
	Protocol        string  `url:"protocol,omitempty"`        //Protocol
	Status          string  `url:"status,omitempty"`          //Status
	TaskID          string  `url:"taskId,omitempty"`          //Task ID
	LastUpdateTime  float64 `url:"lastUpdateTime,omitempty"`  //Last update time
	Limit           float64 `url:"limit,omitempty"`           //Number of resources returned
	Offset          float64 `url:"offset,omitempty"`          //Start index of resources returned (1-based)
	Order           string  `url:"order,omitempty"`           //Order by this field
	SortBy          string  `url:"sortBy,omitempty"`          //Sort by this field
}

type ResponsePathTraceRetrievesAllPreviousPathtracesSummaryV1 struct {
	Response *[]ResponsePathTraceRetrievesAllPreviousPathtracesSummaryV1Response `json:"response,omitempty"` //
	Version  string                                                              `json:"version,omitempty"`  //
}
type ResponsePathTraceRetrievesAllPreviousPathtracesSummaryV1Response struct {
	ControlPath            *bool    `json:"controlPath,omitempty"`            // Control path tracing
	CreateTime             *int     `json:"createTime,omitempty"`             // Timestamp when the Path Trace request was first received
	DestIP                 string   `json:"destIP,omitempty"`                 // IP Address of the destination device
	DestPort               string   `json:"destPort,omitempty"`               // Port on the destination device
	FailureReason          string   `json:"failureReason,omitempty"`          // Reason for failure
	ID                     string   `json:"id,omitempty"`                     // Unique ID for the Path Trace request
	Inclusions             []string `json:"inclusions,omitempty"`             // Subset of {INTERFACE-STATS, QOS-STATS, DEVICE-STATS, PERFORMANCE-STATS, ACL-TRACE}
	LastUpdateTime         *int     `json:"lastUpdateTime,omitempty"`         // Last timestamp when the path trace response was updated
	PeriodicRefresh        *bool    `json:"periodicRefresh,omitempty"`        // Re-run the Path Trace every 30 seconds
	Protocol               string   `json:"protocol,omitempty"`               // One of TCP/UDP or either (null)
	SourceIP               string   `json:"sourceIP,omitempty"`               // IP Address of the source device
	SourcePort             string   `json:"sourcePort,omitempty"`             // Port on the source device
	Status                 string   `json:"status,omitempty"`                 // One of {SUCCESS, INPROGRESS, FAILED, SCHEDULED, PENDING, COMPLETED}
	PreviousFlowAnalysisID string   `json:"previousFlowAnalysisId,omitempty"` // When periodicRefresh is true, this field holds the original Path Trace request ID
}
type ResponsePathTraceInitiateANewPathtraceV1 struct {
	Response *ResponsePathTraceInitiateANewPathtraceV1Response `json:"response,omitempty"` //
	Version  string                                            `json:"version,omitempty"`  //
}
type ResponsePathTraceInitiateANewPathtraceV1Response struct {
	FlowAnalysisID string `json:"flowAnalysisId,omitempty"` //
	TaskID         string `json:"taskId,omitempty"`         //
	URL            string `json:"url,omitempty"`            //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1 struct {
	Response *ResponsePathTraceRetrievesPreviousPathtraceV1Response `json:"response,omitempty"` //
	Version  string                                                 `json:"version,omitempty"`  //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1Response struct {
	DetailedStatus      *ResponsePathTraceRetrievesPreviousPathtraceV1ResponseDetailedStatus        `json:"detailedStatus,omitempty"`      //
	LastUpdate          string                                                                      `json:"lastUpdate,omitempty"`          //
	NetworkElements     *[]ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElements     `json:"networkElements,omitempty"`     //
	NetworkElementsInfo *[]ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfo `json:"networkElementsInfo,omitempty"` //
	Properties          []string                                                                    `json:"properties,omitempty"`          //
	Request             *ResponsePathTraceRetrievesPreviousPathtraceV1ResponseRequest               `json:"request,omitempty"`             //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseDetailedStatus struct {
	ACLTraceCalculation              string `json:"aclTraceCalculation,omitempty"`              //
	ACLTraceCalculationFailureReason string `json:"aclTraceCalculationFailureReason,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElements struct {
	AccuracyList                       *[]ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsAccuracyList           `json:"accuracyList,omitempty"`                       //
	DetailedStatus                     *ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsDetailedStatus           `json:"detailedStatus,omitempty"`                     //
	DeviceStatistics                   *ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsDeviceStatistics         `json:"deviceStatistics,omitempty"`                   //
	DeviceStatsCollection              string                                                                                        `json:"deviceStatsCollection,omitempty"`              //
	DeviceStatsCollectionFailureReason string                                                                                        `json:"deviceStatsCollectionFailureReason,omitempty"` //
	EgressPhysicalInterface            *ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsEgressPhysicalInterface  `json:"egressPhysicalInterface,omitempty"`            //
	EgressVirtualInterface             *ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsEgressVirtualInterface   `json:"egressVirtualInterface,omitempty"`             //
	FlexConnect                        *ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsFlexConnect              `json:"flexConnect,omitempty"`                        //
	ID                                 string                                                                                        `json:"id,omitempty"`                                 //
	IngressPhysicalInterface           *ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsIngressPhysicalInterface `json:"ingressPhysicalInterface,omitempty"`           //
	IngressVirtualInterface            *ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsIngressVirtualInterface  `json:"ingressVirtualInterface,omitempty"`            //
	IP                                 string                                                                                        `json:"ip,omitempty"`                                 //
	LinkInformationSource              string                                                                                        `json:"linkInformationSource,omitempty"`              //
	Name                               string                                                                                        `json:"name,omitempty"`                               //
	PerfMonCollection                  string                                                                                        `json:"perfMonCollection,omitempty"`                  //
	PerfMonCollectionFailureReason     string                                                                                        `json:"perfMonCollectionFailureReason,omitempty"`     //
	PerfMonStatistics                  *[]ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsPerfMonStatistics      `json:"perfMonStatistics,omitempty"`                  //
	Role                               string                                                                                        `json:"role,omitempty"`                               //
	SSID                               string                                                                                        `json:"ssid,omitempty"`                               //
	Tunnels                            []string                                                                                      `json:"tunnels,omitempty"`                            //
	Type                               string                                                                                        `json:"type,omitempty"`                               //
	WLANID                             string                                                                                        `json:"wlanId,omitempty"`                             //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsAccuracyList struct {
	Percent *int   `json:"percent,omitempty"` //
	Reason  string `json:"reason,omitempty"`  //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsDetailedStatus struct {
	ACLTraceCalculation              string `json:"aclTraceCalculation,omitempty"`              //
	ACLTraceCalculationFailureReason string `json:"aclTraceCalculationFailureReason,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsDeviceStatistics struct {
	CPUStatistics    *ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsDeviceStatisticsCPUStatistics    `json:"cpuStatistics,omitempty"`    //
	MemoryStatistics *ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsDeviceStatisticsMemoryStatistics `json:"memoryStatistics,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsDeviceStatisticsCPUStatistics struct {
	FiveMinUsageInPercentage  *float64 `json:"fiveMinUsageInPercentage,omitempty"`  //
	FiveSecsUsageInPercentage *float64 `json:"fiveSecsUsageInPercentage,omitempty"` //
	OneMinUsageInPercentage   *float64 `json:"oneMinUsageInPercentage,omitempty"`   //
	RefreshedAt               *int     `json:"refreshedAt,omitempty"`               //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsDeviceStatisticsMemoryStatistics struct {
	MemoryUsage *int `json:"memoryUsage,omitempty"` //
	RefreshedAt *int `json:"refreshedAt,omitempty"` //
	TotalMemory *int `json:"totalMemory,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsEgressPhysicalInterface struct {
	ACLAnalysis                           *ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsEgressPhysicalInterfaceACLAnalysis         `json:"aclAnalysis,omitempty"`                           //
	ID                                    string                                                                                                          `json:"id,omitempty"`                                    //
	InterfaceStatistics                   *ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsEgressPhysicalInterfaceInterfaceStatistics `json:"interfaceStatistics,omitempty"`                   //
	InterfaceStatsCollection              string                                                                                                          `json:"interfaceStatsCollection,omitempty"`              //
	InterfaceStatsCollectionFailureReason string                                                                                                          `json:"interfaceStatsCollectionFailureReason,omitempty"` //
	Name                                  string                                                                                                          `json:"name,omitempty"`                                  //
	PathOverlayInfo                       *[]ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsEgressPhysicalInterfacePathOverlayInfo   `json:"pathOverlayInfo,omitempty"`                       //
	QosStatistics                         *[]ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsEgressPhysicalInterfaceQosStatistics     `json:"qosStatistics,omitempty"`                         //
	QosStatsCollection                    string                                                                                                          `json:"qosStatsCollection,omitempty"`                    //
	QosStatsCollectionFailureReason       string                                                                                                          `json:"qosStatsCollectionFailureReason,omitempty"`       //
	UsedVLAN                              string                                                                                                          `json:"usedVlan,omitempty"`                              //
	VrfName                               string                                                                                                          `json:"vrfName,omitempty"`                               //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsEgressPhysicalInterfaceACLAnalysis struct {
	ACLName      string                                                                                                                `json:"aclName,omitempty"`      //
	MatchingAces *[]ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsEgressPhysicalInterfaceACLAnalysisMatchingAces `json:"matchingAces,omitempty"` //
	Result       string                                                                                                                `json:"result,omitempty"`       //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsEgressPhysicalInterfaceACLAnalysisMatchingAces struct {
	Ace           string                                                                                                                             `json:"ace,omitempty"`           //
	MatchingPorts *[]ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsEgressPhysicalInterfaceACLAnalysisMatchingAcesMatchingPorts `json:"matchingPorts,omitempty"` //
	Result        string                                                                                                                             `json:"result,omitempty"`        //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsEgressPhysicalInterfaceACLAnalysisMatchingAcesMatchingPorts struct {
	Ports    *[]ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsEgressPhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPorts `json:"ports,omitempty"`    //
	Protocol string                                                                                                                                  `json:"protocol,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsEgressPhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPorts struct {
	DestPorts   []string `json:"destPorts,omitempty"`   //
	SourcePorts []string `json:"sourcePorts,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsEgressPhysicalInterfaceInterfaceStatistics struct {
	AdminStatus        string `json:"adminStatus,omitempty"`        //
	InputPackets       *int   `json:"inputPackets,omitempty"`       //
	InputQueueCount    *int   `json:"inputQueueCount,omitempty"`    //
	InputQueueDrops    *int   `json:"inputQueueDrops,omitempty"`    //
	InputQueueFlushes  *int   `json:"inputQueueFlushes,omitempty"`  //
	InputQueueMaxDepth *int   `json:"inputQueueMaxDepth,omitempty"` //
	InputRatebps       *int   `json:"inputRatebps,omitempty"`       //
	OperationalStatus  string `json:"operationalStatus,omitempty"`  //
	OutputDrop         *int   `json:"outputDrop,omitempty"`         //
	OutputPackets      *int   `json:"outputPackets,omitempty"`      //
	OutputQueueCount   *int   `json:"outputQueueCount,omitempty"`   //
	OutputQueueDepth   *int   `json:"outputQueueDepth,omitempty"`   //
	OutputRatebps      *int   `json:"outputRatebps,omitempty"`      //
	RefreshedAt        *int   `json:"refreshedAt,omitempty"`        //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsEgressPhysicalInterfacePathOverlayInfo struct {
	ControlPlane            string                                                                                                               `json:"controlPlane,omitempty"`            //
	DataPacketEncapsulation string                                                                                                               `json:"dataPacketEncapsulation,omitempty"` //
	DestIP                  string                                                                                                               `json:"destIp,omitempty"`                  //
	DestPort                string                                                                                                               `json:"destPort,omitempty"`                //
	Protocol                string                                                                                                               `json:"protocol,omitempty"`                //
	SourceIP                string                                                                                                               `json:"sourceIp,omitempty"`                //
	SourcePort              string                                                                                                               `json:"sourcePort,omitempty"`              //
	VxlanInfo               *ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsEgressPhysicalInterfacePathOverlayInfoVxlanInfo `json:"vxlanInfo,omitempty"`               //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsEgressPhysicalInterfacePathOverlayInfoVxlanInfo struct {
	Dscp string `json:"dscp,omitempty"` //
	Vnid string `json:"vnid,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsEgressPhysicalInterfaceQosStatistics struct {
	ClassMapName       string `json:"classMapName,omitempty"`       //
	DropRate           *int   `json:"dropRate,omitempty"`           //
	NumBytes           *int   `json:"numBytes,omitempty"`           //
	NumPackets         *int   `json:"numPackets,omitempty"`         //
	OfferedRate        *int   `json:"offeredRate,omitempty"`        //
	QueueBandwidthbps  string `json:"queueBandwidthbps,omitempty"`  //
	QueueDepth         *int   `json:"queueDepth,omitempty"`         //
	QueueNoBufferDrops *int   `json:"queueNoBufferDrops,omitempty"` //
	QueueTotalDrops    *int   `json:"queueTotalDrops,omitempty"`    //
	RefreshedAt        *int   `json:"refreshedAt,omitempty"`        //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsEgressVirtualInterface struct {
	ACLAnalysis                           *ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsEgressVirtualInterfaceACLAnalysis         `json:"aclAnalysis,omitempty"`                           //
	ID                                    string                                                                                                         `json:"id,omitempty"`                                    //
	InterfaceStatistics                   *ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsEgressVirtualInterfaceInterfaceStatistics `json:"interfaceStatistics,omitempty"`                   //
	InterfaceStatsCollection              string                                                                                                         `json:"interfaceStatsCollection,omitempty"`              //
	InterfaceStatsCollectionFailureReason string                                                                                                         `json:"interfaceStatsCollectionFailureReason,omitempty"` //
	Name                                  string                                                                                                         `json:"name,omitempty"`                                  //
	PathOverlayInfo                       *[]ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsEgressVirtualInterfacePathOverlayInfo   `json:"pathOverlayInfo,omitempty"`                       //
	QosStatistics                         *[]ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsEgressVirtualInterfaceQosStatistics     `json:"qosStatistics,omitempty"`                         //
	QosStatsCollection                    string                                                                                                         `json:"qosStatsCollection,omitempty"`                    //
	QosStatsCollectionFailureReason       string                                                                                                         `json:"qosStatsCollectionFailureReason,omitempty"`       //
	UsedVLAN                              string                                                                                                         `json:"usedVlan,omitempty"`                              //
	VrfName                               string                                                                                                         `json:"vrfName,omitempty"`                               //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsEgressVirtualInterfaceACLAnalysis struct {
	ACLName      string                                                                                                               `json:"aclName,omitempty"`      //
	MatchingAces *[]ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsEgressVirtualInterfaceACLAnalysisMatchingAces `json:"matchingAces,omitempty"` //
	Result       string                                                                                                               `json:"result,omitempty"`       //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsEgressVirtualInterfaceACLAnalysisMatchingAces struct {
	Ace           string                                                                                                                            `json:"ace,omitempty"`           //
	MatchingPorts *[]ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsEgressVirtualInterfaceACLAnalysisMatchingAcesMatchingPorts `json:"matchingPorts,omitempty"` //
	Result        string                                                                                                                            `json:"result,omitempty"`        //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsEgressVirtualInterfaceACLAnalysisMatchingAcesMatchingPorts struct {
	Ports    *[]ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsEgressVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPorts `json:"ports,omitempty"`    //
	Protocol string                                                                                                                                 `json:"protocol,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsEgressVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPorts struct {
	DestPorts   []string `json:"destPorts,omitempty"`   //
	SourcePorts []string `json:"sourcePorts,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsEgressVirtualInterfaceInterfaceStatistics struct {
	AdminStatus        string `json:"adminStatus,omitempty"`        //
	InputPackets       *int   `json:"inputPackets,omitempty"`       //
	InputQueueCount    *int   `json:"inputQueueCount,omitempty"`    //
	InputQueueDrops    *int   `json:"inputQueueDrops,omitempty"`    //
	InputQueueFlushes  *int   `json:"inputQueueFlushes,omitempty"`  //
	InputQueueMaxDepth *int   `json:"inputQueueMaxDepth,omitempty"` //
	InputRatebps       *int   `json:"inputRatebps,omitempty"`       //
	OperationalStatus  string `json:"operationalStatus,omitempty"`  //
	OutputDrop         *int   `json:"outputDrop,omitempty"`         //
	OutputPackets      *int   `json:"outputPackets,omitempty"`      //
	OutputQueueCount   *int   `json:"outputQueueCount,omitempty"`   //
	OutputQueueDepth   *int   `json:"outputQueueDepth,omitempty"`   //
	OutputRatebps      *int   `json:"outputRatebps,omitempty"`      //
	RefreshedAt        *int   `json:"refreshedAt,omitempty"`        //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsEgressVirtualInterfacePathOverlayInfo struct {
	ControlPlane            string                                                                                                              `json:"controlPlane,omitempty"`            //
	DataPacketEncapsulation string                                                                                                              `json:"dataPacketEncapsulation,omitempty"` //
	DestIP                  string                                                                                                              `json:"destIp,omitempty"`                  //
	DestPort                string                                                                                                              `json:"destPort,omitempty"`                //
	Protocol                string                                                                                                              `json:"protocol,omitempty"`                //
	SourceIP                string                                                                                                              `json:"sourceIp,omitempty"`                //
	SourcePort              string                                                                                                              `json:"sourcePort,omitempty"`              //
	VxlanInfo               *ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsEgressVirtualInterfacePathOverlayInfoVxlanInfo `json:"vxlanInfo,omitempty"`               //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsEgressVirtualInterfacePathOverlayInfoVxlanInfo struct {
	Dscp string `json:"dscp,omitempty"` //
	Vnid string `json:"vnid,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsEgressVirtualInterfaceQosStatistics struct {
	ClassMapName       string `json:"classMapName,omitempty"`       //
	DropRate           *int   `json:"dropRate,omitempty"`           //
	NumBytes           *int   `json:"numBytes,omitempty"`           //
	NumPackets         *int   `json:"numPackets,omitempty"`         //
	OfferedRate        *int   `json:"offeredRate,omitempty"`        //
	QueueBandwidthbps  string `json:"queueBandwidthbps,omitempty"`  //
	QueueDepth         *int   `json:"queueDepth,omitempty"`         //
	QueueNoBufferDrops *int   `json:"queueNoBufferDrops,omitempty"` //
	QueueTotalDrops    *int   `json:"queueTotalDrops,omitempty"`    //
	RefreshedAt        *int   `json:"refreshedAt,omitempty"`        //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsFlexConnect struct {
	Authentication            string                                                                                             `json:"authentication,omitempty"`            //
	DataSwitching             string                                                                                             `json:"dataSwitching,omitempty"`             //
	EgressACLAnalysis         *ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsFlexConnectEgressACLAnalysis  `json:"egressAclAnalysis,omitempty"`         //
	IngressACLAnalysis        *ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsFlexConnectIngressACLAnalysis `json:"ingressAclAnalysis,omitempty"`        //
	WirelessLanControllerID   string                                                                                             `json:"wirelessLanControllerId,omitempty"`   //
	WirelessLanControllerName string                                                                                             `json:"wirelessLanControllerName,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsFlexConnectEgressACLAnalysis struct {
	ACLName      string                                                                                                          `json:"aclName,omitempty"`      //
	MatchingAces *[]ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsFlexConnectEgressACLAnalysisMatchingAces `json:"matchingAces,omitempty"` //
	Result       string                                                                                                          `json:"result,omitempty"`       //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsFlexConnectEgressACLAnalysisMatchingAces struct {
	Ace           string                                                                                                                       `json:"ace,omitempty"`           //
	MatchingPorts *[]ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsFlexConnectEgressACLAnalysisMatchingAcesMatchingPorts `json:"matchingPorts,omitempty"` //
	Result        string                                                                                                                       `json:"result,omitempty"`        //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsFlexConnectEgressACLAnalysisMatchingAcesMatchingPorts struct {
	Ports    *[]ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsFlexConnectEgressACLAnalysisMatchingAcesMatchingPortsPorts `json:"ports,omitempty"`    //
	Protocol string                                                                                                                            `json:"protocol,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsFlexConnectEgressACLAnalysisMatchingAcesMatchingPortsPorts struct {
	DestPorts   []string `json:"destPorts,omitempty"`   //
	SourcePorts []string `json:"sourcePorts,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsFlexConnectIngressACLAnalysis struct {
	ACLName      string                                                                                                           `json:"aclName,omitempty"`      //
	MatchingAces *[]ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsFlexConnectIngressACLAnalysisMatchingAces `json:"matchingAces,omitempty"` //
	Result       string                                                                                                           `json:"result,omitempty"`       //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsFlexConnectIngressACLAnalysisMatchingAces struct {
	Ace           string                                                                                                                        `json:"ace,omitempty"`           //
	MatchingPorts *[]ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsFlexConnectIngressACLAnalysisMatchingAcesMatchingPorts `json:"matchingPorts,omitempty"` //
	Result        string                                                                                                                        `json:"result,omitempty"`        //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsFlexConnectIngressACLAnalysisMatchingAcesMatchingPorts struct {
	Ports    *[]ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsFlexConnectIngressACLAnalysisMatchingAcesMatchingPortsPorts `json:"ports,omitempty"`    //
	Protocol string                                                                                                                             `json:"protocol,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsFlexConnectIngressACLAnalysisMatchingAcesMatchingPortsPorts struct {
	DestPorts   []string `json:"destPorts,omitempty"`   //
	SourcePorts []string `json:"sourcePorts,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsIngressPhysicalInterface struct {
	ACLAnalysis                           *ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsIngressPhysicalInterfaceACLAnalysis         `json:"aclAnalysis,omitempty"`                           //
	ID                                    string                                                                                                           `json:"id,omitempty"`                                    //
	InterfaceStatistics                   *ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsIngressPhysicalInterfaceInterfaceStatistics `json:"interfaceStatistics,omitempty"`                   //
	InterfaceStatsCollection              string                                                                                                           `json:"interfaceStatsCollection,omitempty"`              //
	InterfaceStatsCollectionFailureReason string                                                                                                           `json:"interfaceStatsCollectionFailureReason,omitempty"` //
	Name                                  string                                                                                                           `json:"name,omitempty"`                                  //
	PathOverlayInfo                       *[]ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsIngressPhysicalInterfacePathOverlayInfo   `json:"pathOverlayInfo,omitempty"`                       //
	QosStatistics                         *[]ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsIngressPhysicalInterfaceQosStatistics     `json:"qosStatistics,omitempty"`                         //
	QosStatsCollection                    string                                                                                                           `json:"qosStatsCollection,omitempty"`                    //
	QosStatsCollectionFailureReason       string                                                                                                           `json:"qosStatsCollectionFailureReason,omitempty"`       //
	UsedVLAN                              string                                                                                                           `json:"usedVlan,omitempty"`                              //
	VrfName                               string                                                                                                           `json:"vrfName,omitempty"`                               //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsIngressPhysicalInterfaceACLAnalysis struct {
	ACLName      string                                                                                                                 `json:"aclName,omitempty"`      //
	MatchingAces *[]ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsIngressPhysicalInterfaceACLAnalysisMatchingAces `json:"matchingAces,omitempty"` //
	Result       string                                                                                                                 `json:"result,omitempty"`       //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsIngressPhysicalInterfaceACLAnalysisMatchingAces struct {
	Ace           string                                                                                                                              `json:"ace,omitempty"`           //
	MatchingPorts *[]ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsIngressPhysicalInterfaceACLAnalysisMatchingAcesMatchingPorts `json:"matchingPorts,omitempty"` //
	Result        string                                                                                                                              `json:"result,omitempty"`        //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsIngressPhysicalInterfaceACLAnalysisMatchingAcesMatchingPorts struct {
	Ports    *[]ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsIngressPhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPorts `json:"ports,omitempty"`    //
	Protocol string                                                                                                                                   `json:"protocol,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsIngressPhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPorts struct {
	DestPorts   []string `json:"destPorts,omitempty"`   //
	SourcePorts []string `json:"sourcePorts,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsIngressPhysicalInterfaceInterfaceStatistics struct {
	AdminStatus        string `json:"adminStatus,omitempty"`        //
	InputPackets       *int   `json:"inputPackets,omitempty"`       //
	InputQueueCount    *int   `json:"inputQueueCount,omitempty"`    //
	InputQueueDrops    *int   `json:"inputQueueDrops,omitempty"`    //
	InputQueueFlushes  *int   `json:"inputQueueFlushes,omitempty"`  //
	InputQueueMaxDepth *int   `json:"inputQueueMaxDepth,omitempty"` //
	InputRatebps       *int   `json:"inputRatebps,omitempty"`       //
	OperationalStatus  string `json:"operationalStatus,omitempty"`  //
	OutputDrop         *int   `json:"outputDrop,omitempty"`         //
	OutputPackets      *int   `json:"outputPackets,omitempty"`      //
	OutputQueueCount   *int   `json:"outputQueueCount,omitempty"`   //
	OutputQueueDepth   *int   `json:"outputQueueDepth,omitempty"`   //
	OutputRatebps      *int   `json:"outputRatebps,omitempty"`      //
	RefreshedAt        *int   `json:"refreshedAt,omitempty"`        //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsIngressPhysicalInterfacePathOverlayInfo struct {
	ControlPlane            string                                                                                                                `json:"controlPlane,omitempty"`            //
	DataPacketEncapsulation string                                                                                                                `json:"dataPacketEncapsulation,omitempty"` //
	DestIP                  string                                                                                                                `json:"destIp,omitempty"`                  //
	DestPort                string                                                                                                                `json:"destPort,omitempty"`                //
	Protocol                string                                                                                                                `json:"protocol,omitempty"`                //
	SourceIP                string                                                                                                                `json:"sourceIp,omitempty"`                //
	SourcePort              string                                                                                                                `json:"sourcePort,omitempty"`              //
	VxlanInfo               *ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsIngressPhysicalInterfacePathOverlayInfoVxlanInfo `json:"vxlanInfo,omitempty"`               //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsIngressPhysicalInterfacePathOverlayInfoVxlanInfo struct {
	Dscp string `json:"dscp,omitempty"` //
	Vnid string `json:"vnid,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsIngressPhysicalInterfaceQosStatistics struct {
	ClassMapName       string `json:"classMapName,omitempty"`       //
	DropRate           *int   `json:"dropRate,omitempty"`           //
	NumBytes           *int   `json:"numBytes,omitempty"`           //
	NumPackets         *int   `json:"numPackets,omitempty"`         //
	OfferedRate        *int   `json:"offeredRate,omitempty"`        //
	QueueBandwidthbps  string `json:"queueBandwidthbps,omitempty"`  //
	QueueDepth         *int   `json:"queueDepth,omitempty"`         //
	QueueNoBufferDrops *int   `json:"queueNoBufferDrops,omitempty"` //
	QueueTotalDrops    *int   `json:"queueTotalDrops,omitempty"`    //
	RefreshedAt        *int   `json:"refreshedAt,omitempty"`        //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsIngressVirtualInterface struct {
	ACLAnalysis                           *ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsIngressVirtualInterfaceACLAnalysis         `json:"aclAnalysis,omitempty"`                           //
	ID                                    string                                                                                                          `json:"id,omitempty"`                                    //
	InterfaceStatistics                   *ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsIngressVirtualInterfaceInterfaceStatistics `json:"interfaceStatistics,omitempty"`                   //
	InterfaceStatsCollection              string                                                                                                          `json:"interfaceStatsCollection,omitempty"`              //
	InterfaceStatsCollectionFailureReason string                                                                                                          `json:"interfaceStatsCollectionFailureReason,omitempty"` //
	Name                                  string                                                                                                          `json:"name,omitempty"`                                  //
	PathOverlayInfo                       *[]ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsIngressVirtualInterfacePathOverlayInfo   `json:"pathOverlayInfo,omitempty"`                       //
	QosStatistics                         *[]ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsIngressVirtualInterfaceQosStatistics     `json:"qosStatistics,omitempty"`                         //
	QosStatsCollection                    string                                                                                                          `json:"qosStatsCollection,omitempty"`                    //
	QosStatsCollectionFailureReason       string                                                                                                          `json:"qosStatsCollectionFailureReason,omitempty"`       //
	UsedVLAN                              string                                                                                                          `json:"usedVlan,omitempty"`                              //
	VrfName                               string                                                                                                          `json:"vrfName,omitempty"`                               //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsIngressVirtualInterfaceACLAnalysis struct {
	ACLName      string                                                                                                                `json:"aclName,omitempty"`      //
	MatchingAces *[]ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsIngressVirtualInterfaceACLAnalysisMatchingAces `json:"matchingAces,omitempty"` //
	Result       string                                                                                                                `json:"result,omitempty"`       //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsIngressVirtualInterfaceACLAnalysisMatchingAces struct {
	Ace           string                                                                                                                             `json:"ace,omitempty"`           //
	MatchingPorts *[]ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsIngressVirtualInterfaceACLAnalysisMatchingAcesMatchingPorts `json:"matchingPorts,omitempty"` //
	Result        string                                                                                                                             `json:"result,omitempty"`        //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsIngressVirtualInterfaceACLAnalysisMatchingAcesMatchingPorts struct {
	Ports    *[]ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsIngressVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPorts `json:"ports,omitempty"`    //
	Protocol string                                                                                                                                  `json:"protocol,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsIngressVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPorts struct {
	DestPorts   []string `json:"destPorts,omitempty"`   //
	SourcePorts []string `json:"sourcePorts,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsIngressVirtualInterfaceInterfaceStatistics struct {
	AdminStatus        string `json:"adminStatus,omitempty"`        //
	InputPackets       *int   `json:"inputPackets,omitempty"`       //
	InputQueueCount    *int   `json:"inputQueueCount,omitempty"`    //
	InputQueueDrops    *int   `json:"inputQueueDrops,omitempty"`    //
	InputQueueFlushes  *int   `json:"inputQueueFlushes,omitempty"`  //
	InputQueueMaxDepth *int   `json:"inputQueueMaxDepth,omitempty"` //
	InputRatebps       *int   `json:"inputRatebps,omitempty"`       //
	OperationalStatus  string `json:"operationalStatus,omitempty"`  //
	OutputDrop         *int   `json:"outputDrop,omitempty"`         //
	OutputPackets      *int   `json:"outputPackets,omitempty"`      //
	OutputQueueCount   *int   `json:"outputQueueCount,omitempty"`   //
	OutputQueueDepth   *int   `json:"outputQueueDepth,omitempty"`   //
	OutputRatebps      *int   `json:"outputRatebps,omitempty"`      //
	RefreshedAt        *int   `json:"refreshedAt,omitempty"`        //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsIngressVirtualInterfacePathOverlayInfo struct {
	ControlPlane            string                                                                                                               `json:"controlPlane,omitempty"`            //
	DataPacketEncapsulation string                                                                                                               `json:"dataPacketEncapsulation,omitempty"` //
	DestIP                  string                                                                                                               `json:"destIp,omitempty"`                  //
	DestPort                string                                                                                                               `json:"destPort,omitempty"`                //
	Protocol                string                                                                                                               `json:"protocol,omitempty"`                //
	SourceIP                string                                                                                                               `json:"sourceIp,omitempty"`                //
	SourcePort              string                                                                                                               `json:"sourcePort,omitempty"`              //
	VxlanInfo               *ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsIngressVirtualInterfacePathOverlayInfoVxlanInfo `json:"vxlanInfo,omitempty"`               //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsIngressVirtualInterfacePathOverlayInfoVxlanInfo struct {
	Dscp string `json:"dscp,omitempty"` //
	Vnid string `json:"vnid,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsIngressVirtualInterfaceQosStatistics struct {
	ClassMapName       string `json:"classMapName,omitempty"`       //
	DropRate           *int   `json:"dropRate,omitempty"`           //
	NumBytes           *int   `json:"numBytes,omitempty"`           //
	NumPackets         *int   `json:"numPackets,omitempty"`         //
	OfferedRate        *int   `json:"offeredRate,omitempty"`        //
	QueueBandwidthbps  string `json:"queueBandwidthbps,omitempty"`  //
	QueueDepth         *int   `json:"queueDepth,omitempty"`         //
	QueueNoBufferDrops *int   `json:"queueNoBufferDrops,omitempty"` //
	QueueTotalDrops    *int   `json:"queueTotalDrops,omitempty"`    //
	RefreshedAt        *int   `json:"refreshedAt,omitempty"`        //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsPerfMonStatistics struct {
	ByteRate             *int     `json:"byteRate,omitempty"`             //
	DestIPAddress        string   `json:"destIpAddress,omitempty"`        //
	DestPort             string   `json:"destPort,omitempty"`             //
	InputInterface       string   `json:"inputInterface,omitempty"`       //
	IPv4DSCP             string   `json:"ipv4DSCP,omitempty"`             //
	IPv4TTL              *int     `json:"ipv4TTL,omitempty"`              //
	OutputInterface      string   `json:"outputInterface,omitempty"`      //
	PacketBytes          *int     `json:"packetBytes,omitempty"`          //
	PacketCount          *int     `json:"packetCount,omitempty"`          //
	PacketLoss           *int     `json:"packetLoss,omitempty"`           //
	PacketLossPercentage *float64 `json:"packetLossPercentage,omitempty"` //
	Protocol             string   `json:"protocol,omitempty"`             //
	RefreshedAt          *int     `json:"refreshedAt,omitempty"`          //
	RtpJitterMax         *int     `json:"rtpJitterMax,omitempty"`         //
	RtpJitterMean        *int     `json:"rtpJitterMean,omitempty"`        //
	RtpJitterMin         *int     `json:"rtpJitterMin,omitempty"`         //
	SourceIPAddress      string   `json:"sourceIpAddress,omitempty"`      //
	SourcePort           string   `json:"sourcePort,omitempty"`           //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfo struct {
	AccuracyList                       *[]ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoAccuracyList          `json:"accuracyList,omitempty"`                       //
	DetailedStatus                     *ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoDetailedStatus          `json:"detailedStatus,omitempty"`                     //
	DeviceStatistics                   *ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoDeviceStatistics        `json:"deviceStatistics,omitempty"`                   //
	DeviceStatsCollection              string                                                                                           `json:"deviceStatsCollection,omitempty"`              //
	DeviceStatsCollectionFailureReason string                                                                                           `json:"deviceStatsCollectionFailureReason,omitempty"` //
	EgressInterface                    *ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoEgressInterface         `json:"egressInterface,omitempty"`                    //
	FlexConnect                        *ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoFlexConnect             `json:"flexConnect,omitempty"`                        //
	ID                                 string                                                                                           `json:"id,omitempty"`                                 //
	IngressInterface                   *ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoIngressInterface        `json:"ingressInterface,omitempty"`                   //
	IP                                 string                                                                                           `json:"ip,omitempty"`                                 //
	LinkInformationSource              string                                                                                           `json:"linkInformationSource,omitempty"`              //
	Name                               string                                                                                           `json:"name,omitempty"`                               //
	PerfMonCollection                  string                                                                                           `json:"perfMonCollection,omitempty"`                  //
	PerfMonCollectionFailureReason     string                                                                                           `json:"perfMonCollectionFailureReason,omitempty"`     //
	PerfMonitorStatistics              *[]ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoPerfMonitorStatistics `json:"perfMonitorStatistics,omitempty"`              //
	Role                               string                                                                                           `json:"role,omitempty"`                               //
	SSID                               string                                                                                           `json:"ssid,omitempty"`                               //
	Tunnels                            []string                                                                                         `json:"tunnels,omitempty"`                            //
	Type                               string                                                                                           `json:"type,omitempty"`                               //
	WLANID                             string                                                                                           `json:"wlanId,omitempty"`                             //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoAccuracyList struct {
	Percent *int   `json:"percent,omitempty"` //
	Reason  string `json:"reason,omitempty"`  //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoDetailedStatus struct {
	ACLTraceCalculation              string `json:"aclTraceCalculation,omitempty"`              //
	ACLTraceCalculationFailureReason string `json:"aclTraceCalculationFailureReason,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoDeviceStatistics struct {
	CPUStatistics    *ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoDeviceStatisticsCPUStatistics    `json:"cpuStatistics,omitempty"`    //
	MemoryStatistics *ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoDeviceStatisticsMemoryStatistics `json:"memoryStatistics,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoDeviceStatisticsCPUStatistics struct {
	FiveMinUsageInPercentage  *float64 `json:"fiveMinUsageInPercentage,omitempty"`  //
	FiveSecsUsageInPercentage *float64 `json:"fiveSecsUsageInPercentage,omitempty"` //
	OneMinUsageInPercentage   *float64 `json:"oneMinUsageInPercentage,omitempty"`   //
	RefreshedAt               *int     `json:"refreshedAt,omitempty"`               //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoDeviceStatisticsMemoryStatistics struct {
	MemoryUsage *int `json:"memoryUsage,omitempty"` //
	RefreshedAt *int `json:"refreshedAt,omitempty"` //
	TotalMemory *int `json:"totalMemory,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoEgressInterface struct {
	PhysicalInterface *ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoEgressInterfacePhysicalInterface  `json:"physicalInterface,omitempty"` //
	VirtualInterface  *[]ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoEgressInterfaceVirtualInterface `json:"virtualInterface,omitempty"`  //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoEgressInterfacePhysicalInterface struct {
	ACLAnalysis                           *ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceACLAnalysis         `json:"aclAnalysis,omitempty"`                           //
	ID                                    string                                                                                                                       `json:"id,omitempty"`                                    //
	InterfaceStatistics                   *ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceInterfaceStatistics `json:"interfaceStatistics,omitempty"`                   //
	InterfaceStatsCollection              string                                                                                                                       `json:"interfaceStatsCollection,omitempty"`              //
	InterfaceStatsCollectionFailureReason string                                                                                                                       `json:"interfaceStatsCollectionFailureReason,omitempty"` //
	Name                                  string                                                                                                                       `json:"name,omitempty"`                                  //
	PathOverlayInfo                       *[]ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoEgressInterfacePhysicalInterfacePathOverlayInfo   `json:"pathOverlayInfo,omitempty"`                       //
	QosStatistics                         *[]ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceQosStatistics     `json:"qosStatistics,omitempty"`                         //
	QosStatsCollection                    string                                                                                                                       `json:"qosStatsCollection,omitempty"`                    //
	QosStatsCollectionFailureReason       string                                                                                                                       `json:"qosStatsCollectionFailureReason,omitempty"`       //
	UsedVLAN                              string                                                                                                                       `json:"usedVlan,omitempty"`                              //
	VrfName                               string                                                                                                                       `json:"vrfName,omitempty"`                               //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceACLAnalysis struct {
	ACLName      string                                                                                                                             `json:"aclName,omitempty"`      //
	MatchingAces *[]ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceACLAnalysisMatchingAces `json:"matchingAces,omitempty"` //
	Result       string                                                                                                                             `json:"result,omitempty"`       //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceACLAnalysisMatchingAces struct {
	Ace           string                                                                                                                                          `json:"ace,omitempty"`           //
	MatchingPorts *[]ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceACLAnalysisMatchingAcesMatchingPorts `json:"matchingPorts,omitempty"` //
	Result        string                                                                                                                                          `json:"result,omitempty"`        //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceACLAnalysisMatchingAcesMatchingPorts struct {
	Ports    *[]ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPorts `json:"ports,omitempty"`    //
	Protocol string                                                                                                                                               `json:"protocol,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPorts struct {
	DestPorts   []string `json:"destPorts,omitempty"`   //
	SourcePorts []string `json:"sourcePorts,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceInterfaceStatistics struct {
	AdminStatus        string `json:"adminStatus,omitempty"`        //
	InputPackets       *int   `json:"inputPackets,omitempty"`       //
	InputQueueCount    *int   `json:"inputQueueCount,omitempty"`    //
	InputQueueDrops    *int   `json:"inputQueueDrops,omitempty"`    //
	InputQueueFlushes  *int   `json:"inputQueueFlushes,omitempty"`  //
	InputQueueMaxDepth *int   `json:"inputQueueMaxDepth,omitempty"` //
	InputRatebps       *int   `json:"inputRatebps,omitempty"`       //
	OperationalStatus  string `json:"operationalStatus,omitempty"`  //
	OutputDrop         *int   `json:"outputDrop,omitempty"`         //
	OutputPackets      *int   `json:"outputPackets,omitempty"`      //
	OutputQueueCount   *int   `json:"outputQueueCount,omitempty"`   //
	OutputQueueDepth   *int   `json:"outputQueueDepth,omitempty"`   //
	OutputRatebps      *int   `json:"outputRatebps,omitempty"`      //
	RefreshedAt        *int   `json:"refreshedAt,omitempty"`        //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoEgressInterfacePhysicalInterfacePathOverlayInfo struct {
	ControlPlane            string                                                                                                                            `json:"controlPlane,omitempty"`            //
	DataPacketEncapsulation string                                                                                                                            `json:"dataPacketEncapsulation,omitempty"` //
	DestIP                  string                                                                                                                            `json:"destIp,omitempty"`                  //
	DestPort                string                                                                                                                            `json:"destPort,omitempty"`                //
	Protocol                string                                                                                                                            `json:"protocol,omitempty"`                //
	SourceIP                string                                                                                                                            `json:"sourceIp,omitempty"`                //
	SourcePort              string                                                                                                                            `json:"sourcePort,omitempty"`              //
	VxlanInfo               *ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoEgressInterfacePhysicalInterfacePathOverlayInfoVxlanInfo `json:"vxlanInfo,omitempty"`               //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoEgressInterfacePhysicalInterfacePathOverlayInfoVxlanInfo struct {
	Dscp string `json:"dscp,omitempty"` //
	Vnid string `json:"vnid,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceQosStatistics struct {
	ClassMapName       string `json:"classMapName,omitempty"`       //
	DropRate           *int   `json:"dropRate,omitempty"`           //
	NumBytes           *int   `json:"numBytes,omitempty"`           //
	NumPackets         *int   `json:"numPackets,omitempty"`         //
	OfferedRate        *int   `json:"offeredRate,omitempty"`        //
	QueueBandwidthbps  string `json:"queueBandwidthbps,omitempty"`  //
	QueueDepth         *int   `json:"queueDepth,omitempty"`         //
	QueueNoBufferDrops *int   `json:"queueNoBufferDrops,omitempty"` //
	QueueTotalDrops    *int   `json:"queueTotalDrops,omitempty"`    //
	RefreshedAt        *int   `json:"refreshedAt,omitempty"`        //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoEgressInterfaceVirtualInterface struct {
	ACLAnalysis                           *ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceACLAnalysis         `json:"aclAnalysis,omitempty"`                           //
	ID                                    string                                                                                                                      `json:"id,omitempty"`                                    //
	InterfaceStatistics                   *ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceInterfaceStatistics `json:"interfaceStatistics,omitempty"`                   //
	InterfaceStatsCollection              string                                                                                                                      `json:"interfaceStatsCollection,omitempty"`              //
	InterfaceStatsCollectionFailureReason string                                                                                                                      `json:"interfaceStatsCollectionFailureReason,omitempty"` //
	Name                                  string                                                                                                                      `json:"name,omitempty"`                                  //
	PathOverlayInfo                       *[]ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoEgressInterfaceVirtualInterfacePathOverlayInfo   `json:"pathOverlayInfo,omitempty"`                       //
	QosStatistics                         *[]ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceQosStatistics     `json:"qosStatistics,omitempty"`                         //
	QosStatsCollection                    string                                                                                                                      `json:"qosStatsCollection,omitempty"`                    //
	QosStatsCollectionFailureReason       string                                                                                                                      `json:"qosStatsCollectionFailureReason,omitempty"`       //
	UsedVLAN                              string                                                                                                                      `json:"usedVlan,omitempty"`                              //
	VrfName                               string                                                                                                                      `json:"vrfName,omitempty"`                               //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceACLAnalysis struct {
	ACLName      string                                                                                                                            `json:"aclName,omitempty"`      //
	MatchingAces *[]ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceACLAnalysisMatchingAces `json:"matchingAces,omitempty"` //
	Result       string                                                                                                                            `json:"result,omitempty"`       //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceACLAnalysisMatchingAces struct {
	Ace           string                                                                                                                                         `json:"ace,omitempty"`           //
	MatchingPorts *[]ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceACLAnalysisMatchingAcesMatchingPorts `json:"matchingPorts,omitempty"` //
	Result        string                                                                                                                                         `json:"result,omitempty"`        //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceACLAnalysisMatchingAcesMatchingPorts struct {
	Ports    *[]ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPorts `json:"ports,omitempty"`    //
	Protocol string                                                                                                                                              `json:"protocol,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPorts struct {
	DestPorts   []string `json:"destPorts,omitempty"`   //
	SourcePorts []string `json:"sourcePorts,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceInterfaceStatistics struct {
	AdminStatus        string `json:"adminStatus,omitempty"`        //
	InputPackets       *int   `json:"inputPackets,omitempty"`       //
	InputQueueCount    *int   `json:"inputQueueCount,omitempty"`    //
	InputQueueDrops    *int   `json:"inputQueueDrops,omitempty"`    //
	InputQueueFlushes  *int   `json:"inputQueueFlushes,omitempty"`  //
	InputQueueMaxDepth *int   `json:"inputQueueMaxDepth,omitempty"` //
	InputRatebps       *int   `json:"inputRatebps,omitempty"`       //
	OperationalStatus  string `json:"operationalStatus,omitempty"`  //
	OutputDrop         *int   `json:"outputDrop,omitempty"`         //
	OutputPackets      *int   `json:"outputPackets,omitempty"`      //
	OutputQueueCount   *int   `json:"outputQueueCount,omitempty"`   //
	OutputQueueDepth   *int   `json:"outputQueueDepth,omitempty"`   //
	OutputRatebps      *int   `json:"outputRatebps,omitempty"`      //
	RefreshedAt        *int   `json:"refreshedAt,omitempty"`        //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoEgressInterfaceVirtualInterfacePathOverlayInfo struct {
	ControlPlane            string                                                                                                                           `json:"controlPlane,omitempty"`            //
	DataPacketEncapsulation string                                                                                                                           `json:"dataPacketEncapsulation,omitempty"` //
	DestIP                  string                                                                                                                           `json:"destIp,omitempty"`                  //
	DestPort                string                                                                                                                           `json:"destPort,omitempty"`                //
	Protocol                string                                                                                                                           `json:"protocol,omitempty"`                //
	SourceIP                string                                                                                                                           `json:"sourceIp,omitempty"`                //
	SourcePort              string                                                                                                                           `json:"sourcePort,omitempty"`              //
	VxlanInfo               *ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoEgressInterfaceVirtualInterfacePathOverlayInfoVxlanInfo `json:"vxlanInfo,omitempty"`               //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoEgressInterfaceVirtualInterfacePathOverlayInfoVxlanInfo struct {
	Dscp string `json:"dscp,omitempty"` //
	Vnid string `json:"vnid,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceQosStatistics struct {
	ClassMapName       string `json:"classMapName,omitempty"`       //
	DropRate           *int   `json:"dropRate,omitempty"`           //
	NumBytes           *int   `json:"numBytes,omitempty"`           //
	NumPackets         *int   `json:"numPackets,omitempty"`         //
	OfferedRate        *int   `json:"offeredRate,omitempty"`        //
	QueueBandwidthbps  string `json:"queueBandwidthbps,omitempty"`  //
	QueueDepth         *int   `json:"queueDepth,omitempty"`         //
	QueueNoBufferDrops *int   `json:"queueNoBufferDrops,omitempty"` //
	QueueTotalDrops    *int   `json:"queueTotalDrops,omitempty"`    //
	RefreshedAt        *int   `json:"refreshedAt,omitempty"`        //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoFlexConnect struct {
	Authentication            string                                                                                                 `json:"authentication,omitempty"`            //
	DataSwitching             string                                                                                                 `json:"dataSwitching,omitempty"`             //
	EgressACLAnalysis         *ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoFlexConnectEgressACLAnalysis  `json:"egressAclAnalysis,omitempty"`         //
	IngressACLAnalysis        *ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoFlexConnectIngressACLAnalysis `json:"ingressAclAnalysis,omitempty"`        //
	WirelessLanControllerID   string                                                                                                 `json:"wirelessLanControllerId,omitempty"`   //
	WirelessLanControllerName string                                                                                                 `json:"wirelessLanControllerName,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoFlexConnectEgressACLAnalysis struct {
	ACLName      string                                                                                                              `json:"aclName,omitempty"`      //
	MatchingAces *[]ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoFlexConnectEgressACLAnalysisMatchingAces `json:"matchingAces,omitempty"` //
	Result       string                                                                                                              `json:"result,omitempty"`       //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoFlexConnectEgressACLAnalysisMatchingAces struct {
	Ace           string                                                                                                                           `json:"ace,omitempty"`           //
	MatchingPorts *[]ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoFlexConnectEgressACLAnalysisMatchingAcesMatchingPorts `json:"matchingPorts,omitempty"` //
	Result        string                                                                                                                           `json:"result,omitempty"`        //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoFlexConnectEgressACLAnalysisMatchingAcesMatchingPorts struct {
	Ports    *[]ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoFlexConnectEgressACLAnalysisMatchingAcesMatchingPortsPorts `json:"ports,omitempty"`    //
	Protocol string                                                                                                                                `json:"protocol,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoFlexConnectEgressACLAnalysisMatchingAcesMatchingPortsPorts struct {
	DestPorts   []string `json:"destPorts,omitempty"`   //
	SourcePorts []string `json:"sourcePorts,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoFlexConnectIngressACLAnalysis struct {
	ACLName      string                                                                                                               `json:"aclName,omitempty"`      //
	MatchingAces *[]ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoFlexConnectIngressACLAnalysisMatchingAces `json:"matchingAces,omitempty"` //
	Result       string                                                                                                               `json:"result,omitempty"`       //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoFlexConnectIngressACLAnalysisMatchingAces struct {
	Ace           string                                                                                                                            `json:"ace,omitempty"`           //
	MatchingPorts *[]ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoFlexConnectIngressACLAnalysisMatchingAcesMatchingPorts `json:"matchingPorts,omitempty"` //
	Result        string                                                                                                                            `json:"result,omitempty"`        //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoFlexConnectIngressACLAnalysisMatchingAcesMatchingPorts struct {
	Ports    *[]ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoFlexConnectIngressACLAnalysisMatchingAcesMatchingPortsPorts `json:"ports,omitempty"`    //
	Protocol string                                                                                                                                 `json:"protocol,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoFlexConnectIngressACLAnalysisMatchingAcesMatchingPortsPorts struct {
	DestPorts   []string `json:"destPorts,omitempty"`   //
	SourcePorts []string `json:"sourcePorts,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoIngressInterface struct {
	PhysicalInterface *ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoIngressInterfacePhysicalInterface  `json:"physicalInterface,omitempty"` //
	VirtualInterface  *[]ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoIngressInterfaceVirtualInterface `json:"virtualInterface,omitempty"`  //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoIngressInterfacePhysicalInterface struct {
	ACLAnalysis                           *ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceACLAnalysis         `json:"aclAnalysis,omitempty"`                           //
	ID                                    string                                                                                                                        `json:"id,omitempty"`                                    //
	InterfaceStatistics                   *ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceInterfaceStatistics `json:"interfaceStatistics,omitempty"`                   //
	InterfaceStatsCollection              string                                                                                                                        `json:"interfaceStatsCollection,omitempty"`              //
	InterfaceStatsCollectionFailureReason string                                                                                                                        `json:"interfaceStatsCollectionFailureReason,omitempty"` //
	Name                                  string                                                                                                                        `json:"name,omitempty"`                                  //
	PathOverlayInfo                       *[]ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoIngressInterfacePhysicalInterfacePathOverlayInfo   `json:"pathOverlayInfo,omitempty"`                       //
	QosStatistics                         *[]ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceQosStatistics     `json:"qosStatistics,omitempty"`                         //
	QosStatsCollection                    string                                                                                                                        `json:"qosStatsCollection,omitempty"`                    //
	QosStatsCollectionFailureReason       string                                                                                                                        `json:"qosStatsCollectionFailureReason,omitempty"`       //
	UsedVLAN                              string                                                                                                                        `json:"usedVlan,omitempty"`                              //
	VrfName                               string                                                                                                                        `json:"vrfName,omitempty"`                               //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceACLAnalysis struct {
	ACLName      string                                                                                                                              `json:"aclName,omitempty"`      //
	MatchingAces *[]ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceACLAnalysisMatchingAces `json:"matchingAces,omitempty"` //
	Result       string                                                                                                                              `json:"result,omitempty"`       //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceACLAnalysisMatchingAces struct {
	Ace           string                                                                                                                                           `json:"ace,omitempty"`           //
	MatchingPorts *[]ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceACLAnalysisMatchingAcesMatchingPorts `json:"matchingPorts,omitempty"` //
	Result        string                                                                                                                                           `json:"result,omitempty"`        //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceACLAnalysisMatchingAcesMatchingPorts struct {
	Ports    *[]ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPorts `json:"ports,omitempty"`    //
	Protocol string                                                                                                                                                `json:"protocol,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPorts struct {
	DestPorts   []string `json:"destPorts,omitempty"`   //
	SourcePorts []string `json:"sourcePorts,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceInterfaceStatistics struct {
	AdminStatus        string `json:"adminStatus,omitempty"`        //
	InputPackets       *int   `json:"inputPackets,omitempty"`       //
	InputQueueCount    *int   `json:"inputQueueCount,omitempty"`    //
	InputQueueDrops    *int   `json:"inputQueueDrops,omitempty"`    //
	InputQueueFlushes  *int   `json:"inputQueueFlushes,omitempty"`  //
	InputQueueMaxDepth *int   `json:"inputQueueMaxDepth,omitempty"` //
	InputRatebps       *int   `json:"inputRatebps,omitempty"`       //
	OperationalStatus  string `json:"operationalStatus,omitempty"`  //
	OutputDrop         *int   `json:"outputDrop,omitempty"`         //
	OutputPackets      *int   `json:"outputPackets,omitempty"`      //
	OutputQueueCount   *int   `json:"outputQueueCount,omitempty"`   //
	OutputQueueDepth   *int   `json:"outputQueueDepth,omitempty"`   //
	OutputRatebps      *int   `json:"outputRatebps,omitempty"`      //
	RefreshedAt        *int   `json:"refreshedAt,omitempty"`        //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoIngressInterfacePhysicalInterfacePathOverlayInfo struct {
	ControlPlane            string                                                                                                                             `json:"controlPlane,omitempty"`            //
	DataPacketEncapsulation string                                                                                                                             `json:"dataPacketEncapsulation,omitempty"` //
	DestIP                  string                                                                                                                             `json:"destIp,omitempty"`                  //
	DestPort                string                                                                                                                             `json:"destPort,omitempty"`                //
	Protocol                string                                                                                                                             `json:"protocol,omitempty"`                //
	SourceIP                string                                                                                                                             `json:"sourceIp,omitempty"`                //
	SourcePort              string                                                                                                                             `json:"sourcePort,omitempty"`              //
	VxlanInfo               *ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoIngressInterfacePhysicalInterfacePathOverlayInfoVxlanInfo `json:"vxlanInfo,omitempty"`               //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoIngressInterfacePhysicalInterfacePathOverlayInfoVxlanInfo struct {
	Dscp string `json:"dscp,omitempty"` //
	Vnid string `json:"vnid,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceQosStatistics struct {
	ClassMapName       string `json:"classMapName,omitempty"`       //
	DropRate           *int   `json:"dropRate,omitempty"`           //
	NumBytes           *int   `json:"numBytes,omitempty"`           //
	NumPackets         *int   `json:"numPackets,omitempty"`         //
	OfferedRate        *int   `json:"offeredRate,omitempty"`        //
	QueueBandwidthbps  string `json:"queueBandwidthbps,omitempty"`  //
	QueueDepth         *int   `json:"queueDepth,omitempty"`         //
	QueueNoBufferDrops *int   `json:"queueNoBufferDrops,omitempty"` //
	QueueTotalDrops    *int   `json:"queueTotalDrops,omitempty"`    //
	RefreshedAt        *int   `json:"refreshedAt,omitempty"`        //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoIngressInterfaceVirtualInterface struct {
	ACLAnalysis                           *ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceACLAnalysis         `json:"aclAnalysis,omitempty"`                           //
	ID                                    string                                                                                                                       `json:"id,omitempty"`                                    //
	InterfaceStatistics                   *ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceInterfaceStatistics `json:"interfaceStatistics,omitempty"`                   //
	InterfaceStatsCollection              string                                                                                                                       `json:"interfaceStatsCollection,omitempty"`              //
	InterfaceStatsCollectionFailureReason string                                                                                                                       `json:"interfaceStatsCollectionFailureReason,omitempty"` //
	Name                                  string                                                                                                                       `json:"name,omitempty"`                                  //
	PathOverlayInfo                       *[]ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoIngressInterfaceVirtualInterfacePathOverlayInfo   `json:"pathOverlayInfo,omitempty"`                       //
	QosStatistics                         *[]ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceQosStatistics     `json:"qosStatistics,omitempty"`                         //
	QosStatsCollection                    string                                                                                                                       `json:"qosStatsCollection,omitempty"`                    //
	QosStatsCollectionFailureReason       string                                                                                                                       `json:"qosStatsCollectionFailureReason,omitempty"`       //
	UsedVLAN                              string                                                                                                                       `json:"usedVlan,omitempty"`                              //
	VrfName                               string                                                                                                                       `json:"vrfName,omitempty"`                               //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceACLAnalysis struct {
	ACLName      string                                                                                                                             `json:"aclName,omitempty"`      //
	MatchingAces *[]ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceACLAnalysisMatchingAces `json:"matchingAces,omitempty"` //
	Result       string                                                                                                                             `json:"result,omitempty"`       //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceACLAnalysisMatchingAces struct {
	Ace           string                                                                                                                                          `json:"ace,omitempty"`           //
	MatchingPorts *[]ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceACLAnalysisMatchingAcesMatchingPorts `json:"matchingPorts,omitempty"` //
	Result        string                                                                                                                                          `json:"result,omitempty"`        //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceACLAnalysisMatchingAcesMatchingPorts struct {
	Ports    *[]ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPorts `json:"ports,omitempty"`    //
	Protocol string                                                                                                                                               `json:"protocol,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPorts struct {
	DestPorts   []string `json:"destPorts,omitempty"`   //
	SourcePorts []string `json:"sourcePorts,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceInterfaceStatistics struct {
	AdminStatus        string `json:"adminStatus,omitempty"`        //
	InputPackets       *int   `json:"inputPackets,omitempty"`       //
	InputQueueCount    *int   `json:"inputQueueCount,omitempty"`    //
	InputQueueDrops    *int   `json:"inputQueueDrops,omitempty"`    //
	InputQueueFlushes  *int   `json:"inputQueueFlushes,omitempty"`  //
	InputQueueMaxDepth *int   `json:"inputQueueMaxDepth,omitempty"` //
	InputRatebps       *int   `json:"inputRatebps,omitempty"`       //
	OperationalStatus  string `json:"operationalStatus,omitempty"`  //
	OutputDrop         *int   `json:"outputDrop,omitempty"`         //
	OutputPackets      *int   `json:"outputPackets,omitempty"`      //
	OutputQueueCount   *int   `json:"outputQueueCount,omitempty"`   //
	OutputQueueDepth   *int   `json:"outputQueueDepth,omitempty"`   //
	OutputRatebps      *int   `json:"outputRatebps,omitempty"`      //
	RefreshedAt        *int   `json:"refreshedAt,omitempty"`        //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoIngressInterfaceVirtualInterfacePathOverlayInfo struct {
	ControlPlane            string                                                                                                                            `json:"controlPlane,omitempty"`            //
	DataPacketEncapsulation string                                                                                                                            `json:"dataPacketEncapsulation,omitempty"` //
	DestIP                  string                                                                                                                            `json:"destIp,omitempty"`                  //
	DestPort                string                                                                                                                            `json:"destPort,omitempty"`                //
	Protocol                string                                                                                                                            `json:"protocol,omitempty"`                //
	SourceIP                string                                                                                                                            `json:"sourceIp,omitempty"`                //
	SourcePort              string                                                                                                                            `json:"sourcePort,omitempty"`              //
	VxlanInfo               *ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoIngressInterfaceVirtualInterfacePathOverlayInfoVxlanInfo `json:"vxlanInfo,omitempty"`               //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoIngressInterfaceVirtualInterfacePathOverlayInfoVxlanInfo struct {
	Dscp string `json:"dscp,omitempty"` //
	Vnid string `json:"vnid,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceQosStatistics struct {
	ClassMapName       string `json:"classMapName,omitempty"`       //
	DropRate           *int   `json:"dropRate,omitempty"`           //
	NumBytes           *int   `json:"numBytes,omitempty"`           //
	NumPackets         *int   `json:"numPackets,omitempty"`         //
	OfferedRate        *int   `json:"offeredRate,omitempty"`        //
	QueueBandwidthbps  string `json:"queueBandwidthbps,omitempty"`  //
	QueueDepth         *int   `json:"queueDepth,omitempty"`         //
	QueueNoBufferDrops *int   `json:"queueNoBufferDrops,omitempty"` //
	QueueTotalDrops    *int   `json:"queueTotalDrops,omitempty"`    //
	RefreshedAt        *int   `json:"refreshedAt,omitempty"`        //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoPerfMonitorStatistics struct {
	ByteRate             *int     `json:"byteRate,omitempty"`             //
	DestIPAddress        string   `json:"destIpAddress,omitempty"`        //
	DestPort             string   `json:"destPort,omitempty"`             //
	InputInterface       string   `json:"inputInterface,omitempty"`       //
	IPv4DSCP             string   `json:"ipv4DSCP,omitempty"`             //
	IPv4TTL              *int     `json:"ipv4TTL,omitempty"`              //
	OutputInterface      string   `json:"outputInterface,omitempty"`      //
	PacketBytes          *int     `json:"packetBytes,omitempty"`          //
	PacketCount          *int     `json:"packetCount,omitempty"`          //
	PacketLoss           *int     `json:"packetLoss,omitempty"`           //
	PacketLossPercentage *float64 `json:"packetLossPercentage,omitempty"` //
	Protocol             string   `json:"protocol,omitempty"`             //
	RefreshedAt          *int     `json:"refreshedAt,omitempty"`          //
	RtpJitterMax         *int     `json:"rtpJitterMax,omitempty"`         //
	RtpJitterMean        *int     `json:"rtpJitterMean,omitempty"`        //
	RtpJitterMin         *int     `json:"rtpJitterMin,omitempty"`         //
	SourceIPAddress      string   `json:"sourceIpAddress,omitempty"`      //
	SourcePort           string   `json:"sourcePort,omitempty"`           //
}
type ResponsePathTraceRetrievesPreviousPathtraceV1ResponseRequest struct {
	ControlPath            *bool    `json:"controlPath,omitempty"`            // Control path tracing
	CreateTime             *int     `json:"createTime,omitempty"`             // Timestamp when the Path Trace request was first received
	DestIP                 string   `json:"destIP,omitempty"`                 // IP Address of the destination device
	DestPort               string   `json:"destPort,omitempty"`               // Port on the destination device
	FailureReason          string   `json:"failureReason,omitempty"`          // Reason for failure
	ID                     string   `json:"id,omitempty"`                     // Unique ID for the Path Trace request
	Inclusions             []string `json:"inclusions,omitempty"`             // Subset of {INTERFACE-STATS, QOS-STATS, DEVICE-STATS, PERFORMANCE-STATS, ACL-TRACE}
	LastUpdateTime         *int     `json:"lastUpdateTime,omitempty"`         // Last timestamp when the path trace response was updated
	PeriodicRefresh        *bool    `json:"periodicRefresh,omitempty"`        // Re-run the Path Trace every 30 seconds
	Protocol               string   `json:"protocol,omitempty"`               // One of TCP/UDP or either (null)
	SourceIP               string   `json:"sourceIP,omitempty"`               // IP Address of the source device
	SourcePort             string   `json:"sourcePort,omitempty"`             // Port on the source device
	Status                 string   `json:"status,omitempty"`                 // One of {SUCCESS, INPROGRESS, FAILED, SCHEDULED, PENDING, COMPLETED}
	PreviousFlowAnalysisID string   `json:"previousFlowAnalysisId,omitempty"` // When periodicRefresh is true, this field holds the original Path Trace request ID
}
type ResponsePathTraceDeletesPathtraceByIDV1 struct {
	Response *ResponsePathTraceDeletesPathtraceByIDV1Response `json:"response,omitempty"` //
	Version  string                                           `json:"version,omitempty"`  //
}
type ResponsePathTraceDeletesPathtraceByIDV1Response struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type RequestPathTraceInitiateANewPathtraceV1 struct {
	ControlPath     *bool    `json:"controlPath,omitempty"`     // Control path tracing
	DestIP          string   `json:"destIP,omitempty"`          // Destination IP address
	DestPort        string   `json:"destPort,omitempty"`        // Destination Port, range: 1-65535
	Inclusions      []string `json:"inclusions,omitempty"`      // Subset of {INTERFACE-STATS, QOS-STATS, DEVICE-STATS, PERFORMANCE-STATS, ACL-TRACE}
	PeriodicRefresh *bool    `json:"periodicRefresh,omitempty"` // Periodic refresh of path for every 30 sec
	Protocol        string   `json:"protocol,omitempty"`        // Protocol - one of [TCP, UDP] - checks both when left blank
	SourceIP        string   `json:"sourceIP,omitempty"`        // Source IP address
	SourcePort      string   `json:"sourcePort,omitempty"`      // Source Port, range: 1-65535
}

//RetrievesAllPreviousPathtracesSummaryV1 Retrieves all previous Pathtraces summary - 55bc-3bf9-4e38-b6ff
/* Returns a summary of all flow analyses stored. Results can be filtered by specified parameters.


@param RetrievesAllPreviousPathtracesSummaryV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-all-previous-pathtraces-summary-v1
*/
func (s *PathTraceService) RetrievesAllPreviousPathtracesSummaryV1(RetrievesAllPreviousPathtracesSummaryV1QueryParams *RetrievesAllPreviousPathtracesSummaryV1QueryParams) (*ResponsePathTraceRetrievesAllPreviousPathtracesSummaryV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/flow-analysis"

	queryString, _ := query.Values(RetrievesAllPreviousPathtracesSummaryV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponsePathTraceRetrievesAllPreviousPathtracesSummaryV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesAllPreviousPathtracesSummaryV1(RetrievesAllPreviousPathtracesSummaryV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesAllPreviousPathtracesSummaryV1")
	}

	result := response.Result().(*ResponsePathTraceRetrievesAllPreviousPathtracesSummaryV1)
	return result, response, err

}

//RetrievesPreviousPathtraceV1 Retrieves previous Pathtrace - 7ab9-a8bd-4f3b-86a4
/* Returns result of a previously requested flow analysis by its Flow Analysis id


@param flowAnalysisID flowAnalysisId path parameter. Flow analysis request id


Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-previous-pathtrace-v1
*/
func (s *PathTraceService) RetrievesPreviousPathtraceV1(flowAnalysisID string) (*ResponsePathTraceRetrievesPreviousPathtraceV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/flow-analysis/{flowAnalysisId}"
	path = strings.Replace(path, "{flowAnalysisId}", fmt.Sprintf("%v", flowAnalysisID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponsePathTraceRetrievesPreviousPathtraceV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesPreviousPathtraceV1(flowAnalysisID)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesPreviousPathtraceV1")
	}

	result := response.Result().(*ResponsePathTraceRetrievesPreviousPathtraceV1)
	return result, response, err

}

//InitiateANewPathtraceV1 Initiate a new Pathtrace - a395-fae6-44ca-899c
/* Initiates a new flow analysis with periodic refresh and stat collection options. Returns a request id and a task id to get results and follow progress.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!initiate-a-new-pathtrace-v1
*/
func (s *PathTraceService) InitiateANewPathtraceV1(requestPathTraceInitiateANewPathtraceV1 *RequestPathTraceInitiateANewPathtraceV1) (*ResponsePathTraceInitiateANewPathtraceV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/flow-analysis"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestPathTraceInitiateANewPathtraceV1).
		SetResult(&ResponsePathTraceInitiateANewPathtraceV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.InitiateANewPathtraceV1(requestPathTraceInitiateANewPathtraceV1)
		}

		return nil, response, fmt.Errorf("error with operation InitiateANewPathtraceV1")
	}

	result := response.Result().(*ResponsePathTraceInitiateANewPathtraceV1)
	return result, response, err

}

//DeletesPathtraceByIDV1 Deletes Pathtrace by Id - 8a9d-2b76-443b-914e
/* Deletes a flow analysis request by its id


@param flowAnalysisID flowAnalysisId path parameter. Flow analysis request id


Documentation Link: https://developer.cisco.com/docs/dna-center/#!deletes-pathtrace-by-id-v1
*/
func (s *PathTraceService) DeletesPathtraceByIDV1(flowAnalysisID string) (*ResponsePathTraceDeletesPathtraceByIDV1, *resty.Response, error) {
	//flowAnalysisID string
	path := "/dna/intent/api/v1/flow-analysis/{flowAnalysisId}"
	path = strings.Replace(path, "{flowAnalysisId}", fmt.Sprintf("%v", flowAnalysisID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponsePathTraceDeletesPathtraceByIDV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeletesPathtraceByIDV1(flowAnalysisID)
		}
		return nil, response, fmt.Errorf("error with operation DeletesPathtraceByIdV1")
	}

	result := response.Result().(*ResponsePathTraceDeletesPathtraceByIDV1)
	return result, response, err

}

// Alias Function
/*
This method acts as an alias for the method `RetrievesPreviousPathtraceV1`
*/
func (s *PathTraceService) RetrievesPreviousPathtrace(flowAnalysisID string) (*ResponsePathTraceRetrievesPreviousPathtraceV1, *resty.Response, error) {
	return s.RetrievesPreviousPathtraceV1(flowAnalysisID)
}

// Alias Function
/*
This method acts as an alias for the method `DeletesPathtraceByIDV1`
*/
func (s *PathTraceService) DeletesPathtraceByID(flowAnalysisID string) (*ResponsePathTraceDeletesPathtraceByIDV1, *resty.Response, error) {
	return s.DeletesPathtraceByIDV1(flowAnalysisID)
}

// Alias Function
/*
This method acts as an alias for the method `RetrievesAllPreviousPathtracesSummaryV1`
*/
func (s *PathTraceService) RetrievesAllPreviousPathtracesSummary(RetrievesAllPreviousPathtracesSummaryV1QueryParams *RetrievesAllPreviousPathtracesSummaryV1QueryParams) (*ResponsePathTraceRetrievesAllPreviousPathtracesSummaryV1, *resty.Response, error) {
	return s.RetrievesAllPreviousPathtracesSummaryV1(RetrievesAllPreviousPathtracesSummaryV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `InitiateANewPathtraceV1`
*/
func (s *PathTraceService) InitiateANewPathtrace(requestPathTraceInitiateANewPathtraceV1 *RequestPathTraceInitiateANewPathtraceV1) (*ResponsePathTraceInitiateANewPathtraceV1, *resty.Response, error) {
	return s.InitiateANewPathtraceV1(requestPathTraceInitiateANewPathtraceV1)
}
