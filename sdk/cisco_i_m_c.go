package catalyst

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
)

type CiscoIMCService service

type ResponseCiscoIMCaddsCiscoIMCConfigurationToACatalystCenterNodeV1 struct {
	Response *ResponseCiscoIMCaddsCiscoIMCConfigurationToACatalystCenterNodeV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // The version of the response
}
type ResponseCiscoIMCaddsCiscoIMCConfigurationToACatalystCenterNodeV1Response struct {
	TaskID string `json:"taskId,omitempty"` // The UUID of the task

	URL string `json:"url,omitempty"` // The path to the API endpoint to GET for information on the task
}
type ResponseCiscoIMCRetrievesCiscoIMCConfigurationsForCatalystCenterNodesV1 struct {
	Response *[]ResponseCiscoIMCRetrievesCiscoIMCConfigurationsForCatalystCenterNodesV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // The version of the response
}
type ResponseCiscoIMCRetrievesCiscoIMCConfigurationsForCatalystCenterNodesV1Response struct {
	ID string `json:"id,omitempty"` // The unique identifier for this Cisco IMC configuration

	NodeID string `json:"nodeId,omitempty"` // The UUID that represents the Catalyst Center node. Its value can be obtained from the `id` attribute of the response of the `/dna/intent/api/v1/nodes-config` API.

	IPAddress string `json:"ipAddress,omitempty"` // IP address of the Cisco IMC

	Username string `json:"username,omitempty"` // Username of the Cisco IMC
}
type ResponseCiscoIMCUpdatesTheCiscoIMCConfigurationForACatalystCenterNodeV1 struct {
	Response *ResponseCiscoIMCUpdatesTheCiscoIMCConfigurationForACatalystCenterNodeV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // The version of the response
}
type ResponseCiscoIMCUpdatesTheCiscoIMCConfigurationForACatalystCenterNodeV1Response struct {
	TaskID string `json:"taskId,omitempty"` // The UUID of the task

	URL string `json:"url,omitempty"` // The path to the API endpoint to GET for information on the task
}
type ResponseCiscoIMCRetrievesTheCiscoIMCConfigurationForACatalystCenterNodeV1 struct {
	Response *ResponseCiscoIMCRetrievesTheCiscoIMCConfigurationForACatalystCenterNodeV1Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // The version of the response
}
type ResponseCiscoIMCRetrievesTheCiscoIMCConfigurationForACatalystCenterNodeV1Response struct {
	ID string `json:"id,omitempty"` // The unique identifier for this Cisco IMC configuration

	NodeID string `json:"nodeId,omitempty"` // The UUID that represents the Catalyst Center node. Its value can be obtained from the `id` attribute of the response of the `/dna/intent/api/v1/nodes-config` API.

	IPAddress string `json:"ipAddress,omitempty"` // IP address of the Cisco IMC

	Username string `json:"username,omitempty"` // Username of the Cisco IMC
}
type RequestCiscoIMCaddsCiscoIMCConfigurationToACatalystCenterNodeV1 struct {
	NodeID string `json:"nodeId,omitempty"` // The UUID that represents the Catalyst Center node. Its value can be obtained from the `id` attribute of the response of the `/dna/intent/api/v1/nodes-config` API.

	IPAddress string `json:"ipAddress,omitempty"` // IP address of the Cisco IMC

	Username string `json:"username,omitempty"` // Username of the Cisco IMC

	Password string `json:"password,omitempty"` // Password of the Cisco IMC
}
type RequestCiscoIMCUpdatesTheCiscoIMCConfigurationForACatalystCenterNodeV1 struct {
	IPAddress string `json:"ipAddress,omitempty"` // IP address of the Cisco IMC

	Username string `json:"username,omitempty"` // Username of the Cisco IMC

	Password string `json:"password,omitempty"` // Password of the Cisco IMC
}

//RetrievesCiscoIMCConfigurationsForCatalystCenterNodesV1 Retrieves Cisco IMC configurations for Catalyst Center nodes - 0082-89c2-45e8-a710
/* This API retrieves the configurations of the Cisco Integrated Management Controller (IMC) that have been added to the Catalyst Center nodes.
The Cisco IMC configuration APIs enable the management of connections between Cisco IMC and Cisco Catalyst Center. By providing the Cisco IMC IP address and credentials to Catalyst Center, Catalyst Center can access and report the health status of hardware components within the Cisco appliance.
More data about the Cisco IMC can be retrieved using the APIs exposed directly by Cisco IMC. Details are available in the Cisco IMC documentation https://www.cisco.com/c/en/us/support/servers-unified-computing/ucs-c-series-integrated-management-controller/series.html#~tab-documents
The Cisco IMC configuration is relevant only for Catalyst Center deployments based on UCS appliances. In cases where Cisco IMC configuration is not supported by the deployment, these APIs will respond with a `404 Not Found` status code.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-cisco-i-m-c-configurations-for-catalyst-center-nodes-v1
*/
func (s *CiscoIMCService) RetrievesCiscoIMCConfigurationsForCatalystCenterNodesV1() (*ResponseCiscoIMCRetrievesCiscoIMCConfigurationsForCatalystCenterNodesV1, *resty.Response, error) {
	path := "/dna/system/api/v1/ciscoImcs"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseCiscoIMCRetrievesCiscoIMCConfigurationsForCatalystCenterNodesV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesCiscoIMCConfigurationsForCatalystCenterNodesV1()
		}
		return nil, response, fmt.Errorf("error with operation RetrievesCiscoIMCConfigurationsForCatalystCenterNodesV1")
	}

	result := response.Result().(*ResponseCiscoIMCRetrievesCiscoIMCConfigurationsForCatalystCenterNodesV1)
	return result, response, err

}

//RetrievesTheCiscoIMCConfigurationForACatalystCenterNodeV1 Retrieves the Cisco IMC configuration for a Catalyst Center node - 859a-6890-4039-b518
/* This API retrieves the Cisco Integrated Management Controller (IMC) configuration for a Catalyst Center node, identified by the specified ID.
The Cisco IMC configuration APIs enable the management of connections between Cisco IMC and Cisco Catalyst Center. By providing the Cisco IMC IP address and credentials to Catalyst Center, Catalyst Center can access and report the health status of hardware components within the Cisco appliance.
More data about the Cisco IMC can be retrieved using the APIs exposed directly by Cisco IMC. Details are available in the Cisco IMC documentation https://www.cisco.com/c/en/us/support/servers-unified-computing/ucs-c-series-integrated-management-controller/series.html#~tab-documents
The Cisco IMC configuration is relevant only for Catalyst Center deployments based on UCS appliances. In cases where Cisco IMC configuration is not supported by the deployment, these APIs will respond with a `404 Not Found` status code.


@param id id path parameter. The unique identifier for this Cisco IMC configuration


Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-cisco-i-m-c-configuration-for-a-catalyst-center-node-v1
*/
func (s *CiscoIMCService) RetrievesTheCiscoIMCConfigurationForACatalystCenterNodeV1(id string) (*ResponseCiscoIMCRetrievesTheCiscoIMCConfigurationForACatalystCenterNodeV1, *resty.Response, error) {
	path := "/dna/system/api/v1/ciscoImcs/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseCiscoIMCRetrievesTheCiscoIMCConfigurationForACatalystCenterNodeV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheCiscoIMCConfigurationForACatalystCenterNodeV1(id)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesTheCiscoIMCConfigurationForACatalystCenterNodeV1")
	}

	result := response.Result().(*ResponseCiscoIMCRetrievesTheCiscoIMCConfigurationForACatalystCenterNodeV1)
	return result, response, err

}

//AddsCiscoIMCConfigurationToACatalystCenterNodeV1 Adds Cisco IMC configuration to a Catalyst Center node - 98a8-a97d-4ba8-b8d1
/* This API adds a Cisco Integrated Management Controller (IMC) configuration to a Cisco Catalyst Center node, identified by its `nodeId`. Obtain the `nodeId` from the `id` attribute in the response of the `/dna/intent/api/v1/nodes-config` API.
The Cisco IMC configuration APIs enable the management of connections between Cisco IMC and Cisco Catalyst Center. By providing the Cisco IMC IP address and credentials to Catalyst Center, Catalyst Center can access and report the health status of hardware components within the Cisco appliance.
More data about the Cisco IMC can be retrieved using the APIs exposed directly by Cisco IMC. Details are available in the Cisco IMC documentation https://www.cisco.com/c/en/us/support/servers-unified-computing/ucs-c-series-integrated-management-controller/series.html#~tab-documents
The Cisco IMC configuration is relevant only for Catalyst Center deployments based on UCS appliances. In cases where Cisco IMC configuration is not supported by the deployment, these APIs will respond with a `404 Not Found` status code.
When Cisco IMC configuration is supported, this API responds with the URL of a diagnostic task.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!adds-cisco-i-m-c-configuration-to-a-catalyst-center-node-v1
*/
func (s *CiscoIMCService) AddsCiscoIMCConfigurationToACatalystCenterNodeV1(requestCiscoIMCAddsCiscoIMCConfigurationToACatalystCenterNodeV1 *RequestCiscoIMCaddsCiscoIMCConfigurationToACatalystCenterNodeV1) (*ResponseCiscoIMCaddsCiscoIMCConfigurationToACatalystCenterNodeV1, *resty.Response, error) {
	path := "/dna/system/api/v1/ciscoImcs"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestCiscoIMCAddsCiscoIMCConfigurationToACatalystCenterNodeV1).
		SetResult(&ResponseCiscoIMCaddsCiscoIMCConfigurationToACatalystCenterNodeV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddsCiscoIMCConfigurationToACatalystCenterNodeV1(requestCiscoIMCAddsCiscoIMCConfigurationToACatalystCenterNodeV1)
		}

		return nil, response, fmt.Errorf("error with operation AddsCiscoIMCConfigurationToACatalystCenterNodeV1")
	}

	result := response.Result().(*ResponseCiscoIMCaddsCiscoIMCConfigurationToACatalystCenterNodeV1)
	return result, response, err

}

//UpdatesTheCiscoIMCConfigurationForACatalystCenterNodeV1 Updates the Cisco IMC configuration for a Catalyst Center node - 6fbc-298f-4879-ad03
/* This API updates the Cisco Integrated Management Controller (IMC) configuration for a Catalyst Center node, identified by the specified ID.
The Cisco IMC configuration APIs enable the management of connections between Cisco IMC and Cisco Catalyst Center. By providing the Cisco IMC IP address and credentials to Catalyst Center, Catalyst Center can access and report the health status of hardware components within the Cisco appliance.
More data about the Cisco IMC can be retrieved using the APIs exposed directly by Cisco IMC. Details are available in the Cisco IMC documentation https://www.cisco.com/c/en/us/support/servers-unified-computing/ucs-c-series-integrated-management-controller/series.html#~tab-documents
The Cisco IMC configuration is relevant only for Catalyst Center deployments based on UCS appliances. In cases where Cisco IMC configuration is not supported by the deployment, these APIs will respond with a `404 Not Found` status code.
When Cisco IMC configuration is supported, this API responds with the URL of a diagnostic task.


@param id id path parameter. The unique identifier for this Cisco IMC configuration

*/
func (s *CiscoIMCService) UpdatesTheCiscoIMCConfigurationForACatalystCenterNodeV1(id string, requestCiscoIMCUpdatesTheCiscoIMCConfigurationForACatalystCenterNodeV1 *RequestCiscoIMCUpdatesTheCiscoIMCConfigurationForACatalystCenterNodeV1) (*ResponseCiscoIMCUpdatesTheCiscoIMCConfigurationForACatalystCenterNodeV1, *resty.Response, error) {
	path := "/dna/system/api/v1/ciscoImcs/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestCiscoIMCUpdatesTheCiscoIMCConfigurationForACatalystCenterNodeV1).
		SetResult(&ResponseCiscoIMCUpdatesTheCiscoIMCConfigurationForACatalystCenterNodeV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdatesTheCiscoIMCConfigurationForACatalystCenterNodeV1(id, requestCiscoIMCUpdatesTheCiscoIMCConfigurationForACatalystCenterNodeV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdatesTheCiscoIMCConfigurationForACatalystCenterNodeV1")
	}

	result := response.Result().(*ResponseCiscoIMCUpdatesTheCiscoIMCConfigurationForACatalystCenterNodeV1)
	return result, response, err

}

//DeletesTheCiscoIMCConfigurationForACatalystCenterNodeV1 Deletes the Cisco IMC configuration for a Catalyst Center node - 0491-cab9-4859-b64d
/* This API removes a specific Cisco Integrated Management Controller (IMC) configuration from a Catalyst Center node using the provided identifier.
The Cisco IMC configuration APIs enable the management of connections between Cisco IMC and Cisco Catalyst Center. By providing the Cisco IMC IP address and credentials to Catalyst Center, Catalyst Center can access and report the health status of hardware components within the Cisco appliance.
More data about the Cisco IMC can be retrieved using the APIs exposed directly by Cisco IMC. Details are available in the Cisco IMC documentation https://www.cisco.com/c/en/us/support/servers-unified-computing/ucs-c-series-integrated-management-controller/series.html#~tab-documents
The Cisco IMC configuration is relevant only for Catalyst Center deployments based on UCS appliances. In cases where Cisco IMC configuration is not supported by the deployment, these APIs will respond with a `404 Not Found` status code.


@param id id path parameter. The unique identifier for this Cisco IMC configuration


Documentation Link: https://developer.cisco.com/docs/dna-center/#!deletes-the-cisco-i-m-c-configuration-for-a-catalyst-center-node-v1
*/
func (s *CiscoIMCService) DeletesTheCiscoIMCConfigurationForACatalystCenterNodeV1(id string) (*resty.Response, error) {
	//id string
	path := "/dna/system/api/v1/ciscoImcs/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeletesTheCiscoIMCConfigurationForACatalystCenterNodeV1(
				id)
		}
		return response, fmt.Errorf("error with operation DeletesTheCiscoIMCConfigurationForACatalystCenterNodeV1")
	}

	return response, err

}

// Alias Function
func (s *CiscoIMCService) AddsCiscoIMCConfigurationToACatalystCenterNode(requestCiscoIMCAddsCiscoIMCConfigurationToACatalystCenterNodeV1 *RequestCiscoIMCaddsCiscoIMCConfigurationToACatalystCenterNodeV1) (*ResponseCiscoIMCaddsCiscoIMCConfigurationToACatalystCenterNodeV1, *resty.Response, error) {
	return s.AddsCiscoIMCConfigurationToACatalystCenterNodeV1(requestCiscoIMCAddsCiscoIMCConfigurationToACatalystCenterNodeV1)
}

// Alias Function
func (s *CiscoIMCService) DeletesTheCiscoIMCConfigurationForACatalystCenterNode(id string) (*resty.Response, error) {
	return s.DeletesTheCiscoIMCConfigurationForACatalystCenterNodeV1(id)
}

// Alias Function
func (s *CiscoIMCService) UpdatesTheCiscoIMCConfigurationForACatalystCenterNode(id string, requestCiscoIMCUpdatesTheCiscoIMCConfigurationForACatalystCenterNodeV1 *RequestCiscoIMCUpdatesTheCiscoIMCConfigurationForACatalystCenterNodeV1) (*ResponseCiscoIMCUpdatesTheCiscoIMCConfigurationForACatalystCenterNodeV1, *resty.Response, error) {
	return s.UpdatesTheCiscoIMCConfigurationForACatalystCenterNodeV1(id, requestCiscoIMCUpdatesTheCiscoIMCConfigurationForACatalystCenterNodeV1)
}

// Alias Function
func (s *CiscoIMCService) RetrievesTheCiscoIMCConfigurationForACatalystCenterNode(id string) (*ResponseCiscoIMCRetrievesTheCiscoIMCConfigurationForACatalystCenterNodeV1, *resty.Response, error) {
	return s.RetrievesTheCiscoIMCConfigurationForACatalystCenterNodeV1(id)
}

// Alias Function
func (s *CiscoIMCService) RetrievesCiscoIMCConfigurationsForCatalystCenterNodes() (*ResponseCiscoIMCRetrievesCiscoIMCConfigurationsForCatalystCenterNodesV1, *resty.Response, error) {
	return s.RetrievesCiscoIMCConfigurationsForCatalystCenterNodesV1()
}
