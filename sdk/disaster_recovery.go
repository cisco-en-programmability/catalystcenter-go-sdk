package catalyst

import (
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
)

type DisasterRecoveryService service

type ResponseDisasterRecoveryDisasterRecoveryOperationalStatusV1 struct {
	Severity       string                                                                 `json:"severity,omitempty"`       // Severity of the DR Event.
	Status         string                                                                 `json:"status,omitempty"`         // Status of the DR Event.
	InitiatedBy    string                                                                 `json:"initiated_by,omitempty"`   // Who initiated this event. Is it a system triggered one or user triggered one.
	IPconfig       *[]ResponseDisasterRecoveryDisasterRecoveryOperationalStatusV1IPconfig `json:"ipconfig,omitempty"`       //
	Tasks          *[]ResponseDisasterRecoveryDisasterRecoveryOperationalStatusV1Tasks    `json:"tasks,omitempty"`          //
	Title          string                                                                 `json:"title,omitempty"`          // DR Event Summary
	Site           string                                                                 `json:"site,omitempty"`           // Site of the DR in which this event occurred.
	StartTimestamp string                                                                 `json:"startTimestamp,omitempty"` // Starting timestamp of the DR event
	Message        string                                                                 `json:"message,omitempty"`        // Detailed Description about the DR event
	EndTimestamp   string                                                                 `json:"endTimestamp,omitempty"`   // End timestamp of the DR event
}
type ResponseDisasterRecoveryDisasterRecoveryOperationalStatusV1IPconfig struct {
	Interface string `json:"interface,omitempty"` // Enterprise or Management interface
	Vip       string `json:"vip,omitempty"`       // Is this interface a Virtual IP address or not. This is true for Site VIP
	IP        string `json:"ip,omitempty"`        // This floating virtual IP address automatically moves to and operates on the site that is currently acting as your network's active site. If the interface is enterprise, then this address manages traffic between your disaster recovery system and your Enterprise network. If the interface is management, then this address manages traffic between your disaster recovery system and your Management network.
}
type ResponseDisasterRecoveryDisasterRecoveryOperationalStatusV1Tasks struct {
	Status         string                                                                      `json:"status,omitempty"`         // Status of the DR event.
	IPconfig       *[]ResponseDisasterRecoveryDisasterRecoveryOperationalStatusV1TasksIPconfig `json:"ipconfig,omitempty"`       //
	Title          string                                                                      `json:"title,omitempty"`          // DR Event Summary
	Site           string                                                                      `json:"site,omitempty"`           // Site of the DR in which this event occured
	StartTimestamp string                                                                      `json:"startTimestamp,omitempty"` // Starting timestamp of the DR event
	Message        string                                                                      `json:"message,omitempty"`        // Detailed description about the DR event
	EndTimestamp   string                                                                      `json:"endTimestamp,omitempty"`   // End timestamp of the DR event
}
type ResponseDisasterRecoveryDisasterRecoveryOperationalStatusV1TasksIPconfig struct {
	Interface string `json:"interface,omitempty"` // Enterprise or Management interface
	Vip       string `json:"vip,omitempty"`       // Is this interface a Virtual IP address or not. This is true for Site VIP
	IP        string `json:"ip,omitempty"`        // This floating virtual IP address automatically moves to and operates on the site that is currently acting as your network's active site. If the interface is enterprise, then this address manages traffic between your disaster recovery system and your Enterprise network. If the interface is management, then this address manages traffic between your disaster recovery system and your Management network.
}
type ResponseDisasterRecoveryDisasterRecoveryStatusV1 struct {
	IPconfig    *[]ResponseDisasterRecoveryDisasterRecoveryStatusV1IPconfig    `json:"ipconfig,omitempty"`     //
	Site        string                                                         `json:"site,omitempty"`         // Site of the disaster recovery system.
	Main        *ResponseDisasterRecoveryDisasterRecoveryStatusV1Main          `json:"main,omitempty"`         //
	Recovery    *ResponseDisasterRecoveryDisasterRecoveryStatusV1Recovery      `json:"recovery,omitempty"`     //
	Witness     *ResponseDisasterRecoveryDisasterRecoveryStatusV1Witness       `json:"witness,omitempty"`      //
	State       string                                                         `json:"state,omitempty"`        // State of the Disaster Recovery System.
	IPsecTunnel *[]ResponseDisasterRecoveryDisasterRecoveryStatusV1IPsecTunnel `json:"ipsec-tunnel,omitempty"` //
}
type ResponseDisasterRecoveryDisasterRecoveryStatusV1IPconfig struct {
	Interface string `json:"interface,omitempty"` // Enterprise or Management interface
	Vip       *bool  `json:"vip,omitempty"`       // Is this interface an Virtual IP address or not. This is true for Global DR VIP
	IP        string `json:"ip,omitempty"`        // This floating virtual IP address automatically moves to and operates on the site that is currently acting as your network's active site. If the interface is enterprise, then this address manages traffic between your disaster recovery system and your Enterprise network. If the interface is management, then this address manages traffic between your disaster recovery system and your Management network.
}
type ResponseDisasterRecoveryDisasterRecoveryStatusV1Main struct {
	IPconfig *[]ResponseDisasterRecoveryDisasterRecoveryStatusV1MainIPconfig `json:"ipconfig,omitempty"` //
	State    string                                                          `json:"state,omitempty"`    // State of the Main Site.
	Nodes    *[]ResponseDisasterRecoveryDisasterRecoveryStatusV1MainNodes    `json:"nodes,omitempty"`    //
}
type ResponseDisasterRecoveryDisasterRecoveryStatusV1MainIPconfig struct {
	Interface string `json:"interface,omitempty"` // Enterprise or Management interface
	Vip       *bool  `json:"vip,omitempty"`       // Is this interface an Virtual IP address or not. This is true for cluster level.
	IP        string `json:"ip,omitempty"`        // This floating virtual IP address automatically moves to and operates on the site that is currently acting as your network's active site. If the interface is enterprise, then this address manages traffic between your disaster recovery system and your Enterprise network. If the interface is management, then this address manages traffic between your disaster recovery system and your Management network.
}
type ResponseDisasterRecoveryDisasterRecoveryStatusV1MainNodes struct {
	Hostname    string                                                                  `json:"hostname,omitempty"`    // Hostname of the node
	State       string                                                                  `json:"state,omitempty"`       // State of the node
	IPaddresses *[]ResponseDisasterRecoveryDisasterRecoveryStatusV1MainNodesIPaddresses `json:"ipaddresses,omitempty"` //
}
type ResponseDisasterRecoveryDisasterRecoveryStatusV1MainNodesIPaddresses struct {
	Interface string `json:"interface,omitempty"` // Enterprise or Management interface
	Vip       *bool  `json:"vip,omitempty"`       // Is this interface a Virtual IP address or not. This is false for node level.
	IP        string `json:"ip,omitempty"`        // Node IP address
}
type ResponseDisasterRecoveryDisasterRecoveryStatusV1Recovery struct {
	IPconfig *[]ResponseDisasterRecoveryDisasterRecoveryStatusV1RecoveryIPconfig `json:"ipconfig,omitempty"` //
	State    string                                                              `json:"state,omitempty"`    // State of the Recovery site
	Nodes    *[]ResponseDisasterRecoveryDisasterRecoveryStatusV1RecoveryNodes    `json:"nodes,omitempty"`    //
}
type ResponseDisasterRecoveryDisasterRecoveryStatusV1RecoveryIPconfig struct {
	Interface string `json:"interface,omitempty"` // Enterprise or Management interface
	Vip       *bool  `json:"vip,omitempty"`       // Is this interface an Virtual IP address or not. This is true for cluster level.
	IP        string `json:"ip,omitempty"`        // This floating virtual IP address automatically moves to and operates on the site that is currently acting as your network's active site. If the interface is enterprise, then this address manages traffic between your disaster recovery system and your Enterprise network. If the interface is management, then this address manages traffic between your disaster recovery system and your Management network.
}
type ResponseDisasterRecoveryDisasterRecoveryStatusV1RecoveryNodes struct {
	Hostname string                                                                   `json:"hostname,omitempty"` // Hostname of the node
	State    string                                                                   `json:"state,omitempty"`    // State of the node
	IPconfig *[]ResponseDisasterRecoveryDisasterRecoveryStatusV1RecoveryNodesIPconfig `json:"ipconfig,omitempty"` //
}
type ResponseDisasterRecoveryDisasterRecoveryStatusV1RecoveryNodesIPconfig struct {
	Interface string `json:"interface,omitempty"` // Enterprise or Management interface
	Vip       *bool  `json:"vip,omitempty"`       // Is this interface a Virtual IP Address or not. This is false for node level.
	IP        string `json:"ip,omitempty"`        // Node IP Address
}
type ResponseDisasterRecoveryDisasterRecoveryStatusV1Witness struct {
	IPconfig *[]ResponseDisasterRecoveryDisasterRecoveryStatusV1WitnessIPconfig `json:"ipconfig,omitempty"` //
	State    string                                                             `json:"state,omitempty"`    // State of the Witness Site
	Nodes    *[]ResponseDisasterRecoveryDisasterRecoveryStatusV1WitnessNodes    `json:"nodes,omitempty"`    //
}
type ResponseDisasterRecoveryDisasterRecoveryStatusV1WitnessIPconfig struct {
	Interface string `json:"interface,omitempty"` // Enterprise or Management interface
	Vip       *bool  `json:"vip,omitempty"`       // Is this interface an Virtual IP address or not. This is false for witness.
	IP        string `json:"ip,omitempty"`        // In case of witness, this is only an IP.
}
type ResponseDisasterRecoveryDisasterRecoveryStatusV1WitnessNodes struct {
	Hostname string                                                                  `json:"hostname,omitempty"` // Hostname of the witness node
	State    string                                                                  `json:"state,omitempty"`    // State of the node
	IPconfig *[]ResponseDisasterRecoveryDisasterRecoveryStatusV1WitnessNodesIPconfig `json:"ipconfig,omitempty"` //
}
type ResponseDisasterRecoveryDisasterRecoveryStatusV1WitnessNodesIPconfig struct {
	Interface string `json:"interface,omitempty"` // Enterprise or Management interface
	Vip       *bool  `json:"vip,omitempty"`       // Is this interface an Virtual IP address or not. This is false for Witness
	IP        string `json:"ip,omitempty"`        // In case of witness, this is only an IP
}
type ResponseDisasterRecoveryDisasterRecoveryStatusV1IPsecTunnel struct {
	SideA  string `json:"side_a,omitempty"` // A Side of the IPSec Tunnel
	SideB  string `json:"side_b,omitempty"` // Other Side of the IPSec Tunnel
	Status string `json:"status,omitempty"` // Status of this IPSec Tunnel
}

//DisasterRecoveryOperationalStatusV1 Disaster Recovery Operational Status - b89c-dbd5-45da-a6e5
/* Returns the status of Disaster Recovery operation performed on the system.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!disaster-recovery-operational-status-v1
*/
func (s *DisasterRecoveryService) DisasterRecoveryOperationalStatusV1() (*ResponseDisasterRecoveryDisasterRecoveryOperationalStatusV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/disasterrecovery/system/operationstatus"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDisasterRecoveryDisasterRecoveryOperationalStatusV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DisasterRecoveryOperationalStatusV1()
		}
		return nil, response, fmt.Errorf("error with operation DisasterRecoveryOperationalStatusV1")
	}

	result := response.Result().(*ResponseDisasterRecoveryDisasterRecoveryOperationalStatusV1)
	return result, response, err

}

//DisasterRecoveryStatusV1 Disaster Recovery Status - 0b83-9ba0-493a-bf2e
/* Detailed and Summarized status of DR components (Active, Standby and Witness system's health).



Documentation Link: https://developer.cisco.com/docs/dna-center/#!disaster-recovery-status-v1
*/
func (s *DisasterRecoveryService) DisasterRecoveryStatusV1() (*ResponseDisasterRecoveryDisasterRecoveryStatusV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/disasterrecovery/system/status"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDisasterRecoveryDisasterRecoveryStatusV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DisasterRecoveryStatusV1()
		}
		return nil, response, fmt.Errorf("error with operation DisasterRecoveryStatusV1")
	}

	result := response.Result().(*ResponseDisasterRecoveryDisasterRecoveryStatusV1)
	return result, response, err

}

// Alias Function
func (s *DisasterRecoveryService) DisasterRecoveryOperationalStatus() (*ResponseDisasterRecoveryDisasterRecoveryOperationalStatusV1, *resty.Response, error) {
	return s.DisasterRecoveryOperationalStatusV1()
}

// Alias Function
func (s *DisasterRecoveryService) DisasterRecoveryStatus() (*ResponseDisasterRecoveryDisasterRecoveryStatusV1, *resty.Response, error) {
	return s.DisasterRecoveryStatusV1()
}
