package catalyst

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type TagService service

type GetTagV1QueryParams struct {
	Name                     string  `url:"name,omitempty"`                      //Tag name is mandatory when filter operation is used.
	AdditionalInfonameSpace  string  `url:"additionalInfo.nameSpace,omitempty"`  //nameSpace
	AdditionalInfoattributes string  `url:"additionalInfo.attributes,omitempty"` //attributeName
	Level                    string  `url:"level,omitempty"`                     //levelArg
	Offset                   float64 `url:"offset,omitempty"`                    //offset
	Limit                    float64 `url:"limit,omitempty"`                     //limit
	Size                     string  `url:"size,omitempty"`                      //size in kilobytes(KB)
	Field                    string  `url:"field,omitempty"`                     //Available field names are :'name,id,parentId,type,additionalInfo.nameSpace,additionalInfo.attributes'
	SortBy                   string  `url:"sortBy,omitempty"`                    //Only supported attribute is name. SortyBy is mandatory when order is used.
	Order                    string  `url:"order,omitempty"`                     //Available values are asc and des
	SystemTag                string  `url:"systemTag,omitempty"`                 //systemTag
}
type GetTagCountV1QueryParams struct {
	Name          string `url:"name,omitempty"`          //tagName
	NameSpace     string `url:"nameSpace,omitempty"`     //nameSpace
	AttributeName string `url:"attributeName,omitempty"` //attributeName
	Size          string `url:"size,omitempty"`          //size in kilobytes(KB)
	SystemTag     string `url:"systemTag,omitempty"`     //systemTag
}
type GetTagMembersByIDV1QueryParams struct {
	MemberType            string  `url:"memberType,omitempty"`            //Entity type of the member. Possible values can be retrieved by using /tag/member/type API
	Offset                float64 `url:"offset,omitempty"`                //Used for pagination. It indicates the starting row number out of available member records
	Limit                 float64 `url:"limit,omitempty"`                 //Used to Number of maximum members to return in the result
	MemberAssociationType string  `url:"memberAssociationType,omitempty"` //Indicates how the member is associated with the tag. Possible values and description. 1) DYNAMIC : The member is associated to the tag through rules. 2) STATIC – The member is associated to the tag manually. 3) MIXED – The member is associated manually and also satisfies the rule defined for the tag
	Level                 string  `url:"level,omitempty"`                 //level
}
type GetTagMemberCountV1QueryParams struct {
	MemberType            string `url:"memberType,omitempty"`            //memberType
	MemberAssociationType string `url:"memberAssociationType,omitempty"` //memberAssociationType
}
type RetrieveTagsAssociatedWithTheInterfacesV1QueryParams struct {
	Offset float64 `url:"offset,omitempty"` //The first record to show for this page; the first record is numbered 1. minimum: 1
	Limit  float64 `url:"limit,omitempty"`  //The number of records to show for this page. minimum: 1, maximum: 500
}
type RetrieveTagsAssociatedWithNetworkDevicesV1QueryParams struct {
	Offset float64 `url:"offset,omitempty"` //The first record to show for this page; the first record is numbered 1. minimum: 1
	Limit  float64 `url:"limit,omitempty"`  //The number of records to show for this page. minimum: 1, maximum: 500
}

type ResponseTagUpdateTagV1 struct {
	Version  string                          `json:"version,omitempty"`  //
	Response *ResponseTagUpdateTagV1Response `json:"response,omitempty"` //
}
type ResponseTagUpdateTagV1Response struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseTagGetTagV1 struct {
	Version  string                         `json:"version,omitempty"`  //
	Response *[]ResponseTagGetTagV1Response `json:"response,omitempty"` //
}
type ResponseTagGetTagV1Response struct {
	SystemTag        *bool                                      `json:"systemTag,omitempty"`        //
	Description      string                                     `json:"description,omitempty"`      //
	DynamicRules     *[]ResponseTagGetTagV1ResponseDynamicRules `json:"dynamicRules,omitempty"`     //
	Name             string                                     `json:"name,omitempty"`             //
	ID               string                                     `json:"id,omitempty"`               //
	InstanceTenantID string                                     `json:"instanceTenantId,omitempty"` //
}
type ResponseTagGetTagV1ResponseDynamicRules struct {
	MemberType string                                        `json:"memberType,omitempty"` //
	Rules      *ResponseTagGetTagV1ResponseDynamicRulesRules `json:"rules,omitempty"`      //
}
type ResponseTagGetTagV1ResponseDynamicRulesRules struct {
	Values    []string                                             `json:"values,omitempty"`    //
	Items     *[]ResponseTagGetTagV1ResponseDynamicRulesRulesItems `json:"items,omitempty"`     //
	Operation string                                               `json:"operation,omitempty"` //
	Name      string                                               `json:"name,omitempty"`      //
	Value     string                                               `json:"value,omitempty"`     //
}
type ResponseTagGetTagV1ResponseDynamicRulesRulesItems interface{}
type ResponseTagCreateTagV1 struct {
	Version  string                          `json:"version,omitempty"`  //
	Response *ResponseTagCreateTagV1Response `json:"response,omitempty"` //
}
type ResponseTagCreateTagV1Response struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseTagGetTagCountV1 struct {
	Version  string `json:"version,omitempty"`  //
	Response *int   `json:"response,omitempty"` //
}
type ResponseTagUpdateTagMembershipV1 struct {
	Version  string                                    `json:"version,omitempty"`  //
	Response *ResponseTagUpdateTagMembershipV1Response `json:"response,omitempty"` //
}
type ResponseTagUpdateTagMembershipV1Response struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseTagGetTagResourceTypesV1 struct {
	Version  string   `json:"version,omitempty"`  //
	Response []string `json:"response,omitempty"` //
}
type ResponseTagDeleteTagV1 struct {
	Version  string                          `json:"version,omitempty"`  //
	Response *ResponseTagDeleteTagV1Response `json:"response,omitempty"` //
}
type ResponseTagDeleteTagV1Response struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseTagGetTagByIDV1 struct {
	Version  string                           `json:"version,omitempty"`  //
	Response *ResponseTagGetTagByIDV1Response `json:"response,omitempty"` //
}
type ResponseTagGetTagByIDV1Response struct {
	SystemTag        *bool                                          `json:"systemTag,omitempty"`        //
	Description      string                                         `json:"description,omitempty"`      //
	DynamicRules     *[]ResponseTagGetTagByIDV1ResponseDynamicRules `json:"dynamicRules,omitempty"`     //
	Name             string                                         `json:"name,omitempty"`             //
	ID               string                                         `json:"id,omitempty"`               //
	InstanceTenantID string                                         `json:"instanceTenantId,omitempty"` //
}
type ResponseTagGetTagByIDV1ResponseDynamicRules struct {
	MemberType string                                            `json:"memberType,omitempty"` //
	Rules      *ResponseTagGetTagByIDV1ResponseDynamicRulesRules `json:"rules,omitempty"`      //
}
type ResponseTagGetTagByIDV1ResponseDynamicRulesRules struct {
	Values    []string                                                 `json:"values,omitempty"`    //
	Items     *[]ResponseTagGetTagByIDV1ResponseDynamicRulesRulesItems `json:"items,omitempty"`     //
	Operation string                                                   `json:"operation,omitempty"` //
	Name      string                                                   `json:"name,omitempty"`      //
	Value     string                                                   `json:"value,omitempty"`     //
}
type ResponseTagGetTagByIDV1ResponseDynamicRulesRulesItems interface{}
type ResponseTagGetTagMembersByIDV1 struct {
	Version  string                                    `json:"version,omitempty"`  //
	Response *[]ResponseTagGetTagMembersByIDV1Response `json:"response,omitempty"` //
}
type ResponseTagGetTagMembersByIDV1Response struct {
	InstanceUUID string `json:"instanceUuid,omitempty"` //
}
type ResponseTagAddMembersToTheTagV1 struct {
	Version  string                                   `json:"version,omitempty"`  //
	Response *ResponseTagAddMembersToTheTagV1Response `json:"response,omitempty"` //
}
type ResponseTagAddMembersToTheTagV1Response struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseTagGetTagMemberCountV1 struct {
	Version  string `json:"version,omitempty"`  //
	Response *int   `json:"response,omitempty"` //
}
type ResponseTagRemoveTagMemberV1 struct {
	Version  string                                `json:"version,omitempty"`  //
	Response *ResponseTagRemoveTagMemberV1Response `json:"response,omitempty"` //
}
type ResponseTagRemoveTagMemberV1Response struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseTagRetrieveTagsAssociatedWithTheInterfacesV1 struct {
	Response *[]ResponseTagRetrieveTagsAssociatedWithTheInterfacesV1Response `json:"response,omitempty"` //
	Version  string                                                          `json:"version,omitempty"`  // The version of the response.
}
type ResponseTagRetrieveTagsAssociatedWithTheInterfacesV1Response struct {
	ID   string                                                              `json:"id,omitempty"`   // Id of the member (network device or interface)
	Tags *[]ResponseTagRetrieveTagsAssociatedWithTheInterfacesV1ResponseTags `json:"tags,omitempty"` //
}
type ResponseTagRetrieveTagsAssociatedWithTheInterfacesV1ResponseTags struct {
	ID   string `json:"id,omitempty"`   // Tag id
	Name string `json:"name,omitempty"` // Tag name
}
type ResponseTagRetrieveTheCountOfInterfacesThatAreAssociatedWithAtLeastOneTagV1 struct {
	Response *ResponseTagRetrieveTheCountOfInterfacesThatAreAssociatedWithAtLeastOneTagV1Response `json:"response,omitempty"` //
	Version  string                                                                               `json:"version,omitempty"`  // The version of the response.
}
type ResponseTagRetrieveTheCountOfInterfacesThatAreAssociatedWithAtLeastOneTagV1Response struct {
	Count *int `json:"count,omitempty"` // The reported count.
}
type ResponseTagQueryTheTagsAssociatedWithInterfacesV1 struct {
	Response *[]ResponseTagQueryTheTagsAssociatedWithInterfacesV1Response `json:"response,omitempty"` //
	Version  string                                                       `json:"version,omitempty"`  // The version of the response.
}
type ResponseTagQueryTheTagsAssociatedWithInterfacesV1Response struct {
	ID   string                                                           `json:"id,omitempty"`   // Id of the member (network device or interface)
	Tags *[]ResponseTagQueryTheTagsAssociatedWithInterfacesV1ResponseTags `json:"tags,omitempty"` //
}
type ResponseTagQueryTheTagsAssociatedWithInterfacesV1ResponseTags struct {
	ID   string `json:"id,omitempty"`   // Tag id
	Name string `json:"name,omitempty"` // Tag name
}
type ResponseTagRetrieveTagsAssociatedWithNetworkDevicesV1 struct {
	Response *[]ResponseTagRetrieveTagsAssociatedWithNetworkDevicesV1Response `json:"response,omitempty"` //
	Version  string                                                           `json:"version,omitempty"`  // The version of the response.
}
type ResponseTagRetrieveTagsAssociatedWithNetworkDevicesV1Response struct {
	ID   string                                                               `json:"id,omitempty"`   // Id of the member (network device or interface)
	Tags *[]ResponseTagRetrieveTagsAssociatedWithNetworkDevicesV1ResponseTags `json:"tags,omitempty"` //
}
type ResponseTagRetrieveTagsAssociatedWithNetworkDevicesV1ResponseTags struct {
	ID   string `json:"id,omitempty"`   // Tag id
	Name string `json:"name,omitempty"` // Tag name
}
type ResponseTagRetrieveTheCountOfNetworkDevicesThatAreAssociatedWithAtLeastOneTagV1 struct {
	Response *ResponseTagRetrieveTheCountOfNetworkDevicesThatAreAssociatedWithAtLeastOneTagV1Response `json:"response,omitempty"` //
	Version  string                                                                                   `json:"version,omitempty"`  // The version of the response.
}
type ResponseTagRetrieveTheCountOfNetworkDevicesThatAreAssociatedWithAtLeastOneTagV1Response struct {
	Count *int `json:"count,omitempty"` // The reported count.
}
type ResponseTagQueryTheTagsAssociatedWithNetworkDevicesV1 struct {
	Response *[]ResponseTagQueryTheTagsAssociatedWithNetworkDevicesV1Response `json:"response,omitempty"` //
	Version  string                                                           `json:"version,omitempty"`  // The version of the response.
}
type ResponseTagQueryTheTagsAssociatedWithNetworkDevicesV1Response struct {
	ID   string                                                               `json:"id,omitempty"`   // Id of the member (network device or interface)
	Tags *[]ResponseTagQueryTheTagsAssociatedWithNetworkDevicesV1ResponseTags `json:"tags,omitempty"` //
}
type ResponseTagQueryTheTagsAssociatedWithNetworkDevicesV1ResponseTags struct {
	ID   string `json:"id,omitempty"`   // Tag id
	Name string `json:"name,omitempty"` // Tag name
}
type RequestTagUpdateTagV1 struct {
	SystemTag        *bool                                `json:"systemTag,omitempty"`        // true for system created tags, false for user defined tags
	Description      string                               `json:"description,omitempty"`      // description of the tag.
	DynamicRules     *[]RequestTagUpdateTagV1DynamicRules `json:"dynamicRules,omitempty"`     //
	Name             string                               `json:"name,omitempty"`             // name of the tag.
	ID               string                               `json:"id,omitempty"`               // mandatory instanceUuid of the tag that needs to be updated.
	InstanceTenantID string                               `json:"instanceTenantId,omitempty"` // instanceTenantId generated for the tag.
}
type RequestTagUpdateTagV1DynamicRules struct {
	MemberType string                                  `json:"memberType,omitempty"` // memberType of the tag (e.g. networkdevice, interface)
	Rules      *RequestTagUpdateTagV1DynamicRulesRules `json:"rules,omitempty"`      //
}
type RequestTagUpdateTagV1DynamicRulesRules struct {
	Values    []string    `json:"values,omitempty"`    // values of the parameter,Only one of the value or values can be used for the given parameter. (for managementIpAddress e.g. ["10.197.124.90","10.197.124.91"])
	Items     interface{} `json:"items,omitempty"`     // items details,multiple rules can be defined by items(e.g. "items": [{"operation": "ILIKE", "name": "managementIpAddress", "value": "%10%"}, {"operation": "ILIKE", "name": "hostname", "value": "%NA%"} ])
	Operation string      `json:"operation,omitempty"` // opeartion used in the rules (e.g. OR,IN,EQ,LIKE,ILIKE,AND)
	Name      string      `json:"name,omitempty"`      // name of the parameter (e.g. for interface:portName,adminStatus,speed,status,description. for networkdevice:family,series,hostname,managementIpAddress,groupNameHierarchy,softwareVersion)
	Value     string      `json:"value,omitempty"`     // value of the parameter (e.g. for portName:1/0/1,for adminStatus,status:up/down, for speed: any integer value, for description: any valid string, for family:switches, for series:C3650, for managementIpAddress:10.197.124.90, groupNameHierarchy:Global, softwareVersion: 16.9.1)
}

type RequestTagCreateTagV1 struct {
	SystemTag        *bool                                `json:"systemTag,omitempty"`        // true for system created tags, false for user defined tags
	Description      string                               `json:"description,omitempty"`      // description of the tag.
	DynamicRules     *[]RequestTagCreateTagV1DynamicRules `json:"dynamicRules,omitempty"`     //
	Name             string                               `json:"name,omitempty"`             // name of the tag.
	ID               string                               `json:"id,omitempty"`               // instanceUuid generated for the tag.
	InstanceTenantID string                               `json:"instanceTenantId,omitempty"` // instanceTenantId generated for the tag.
}
type RequestTagCreateTagV1DynamicRules struct {
	MemberType string                                  `json:"memberType,omitempty"` // memberType of the tag (e.g. networkdevice, interface)
	Rules      *RequestTagCreateTagV1DynamicRulesRules `json:"rules,omitempty"`      //
}
type RequestTagCreateTagV1DynamicRulesRules struct {
	Values    []string    `json:"values,omitempty"`    // values of the parameter,Only one of the value or values can be used for the given parameter. (for managementIpAddress e.g. ["10.197.124.90","10.197.124.91"])
	Items     interface{} `json:"items,omitempty"`     // items details,multiple rules can be defined by items(e.g. "items": [{"operation": "ILIKE", "name": "managementIpAddress", "value": "%10%"}, {"operation": "ILIKE", "name": "hostname", "value": "%NA%"} ])
	Operation string      `json:"operation,omitempty"` // opeartion used in the rules (e.g. OR,IN,EQ,LIKE,ILIKE,AND)
	Name      string      `json:"name,omitempty"`      // name of the parameter (e.g. for interface:portName,adminStatus,speed,status,description. for networkdevice:family,series,hostname,managementIpAddress,groupNameHierarchy,softwareVersion)
	Value     string      `json:"value,omitempty"`     // value of the parameter (e.g. for portName:1/0/1,for adminStatus,status:up/down, for speed: any integer value, for description: any valid string, for family:switches, for series:C3650, for managementIpAddress:10.197.124.90, groupNameHierarchy:Global, softwareVersion: 16.9.1)
}

type RequestTagUpdateTagMembershipV1 struct {
	MemberToTags map[string][]string `json:"memberToTags,omitempty"` //
	MemberType   string              `json:"memberType,omitempty"`   //
}
type RequestTagUpdateTagMembershipV1MemberToTags struct {
	Key []string `json:"key,omitempty"` //
}
type RequestTagAddMembersToTheTagV1 map[string][]string
type RequestTagQueryTheTagsAssociatedWithInterfacesV1 struct {
	IDs []string `json:"ids,omitempty"` // List of member ids (network device or interface), maximum 500 ids can be passed.
}
type RequestTagQueryTheTagsAssociatedWithNetworkDevicesV1 struct {
	IDs []string `json:"ids,omitempty"` // List of member ids (network device or interface), maximum 500 ids can be passed.
}

//GetTagV1 Get Tag - ee9a-ab01-487a-8896
/* Returns the tags for given filter criteria


@param GetTagV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-tag-v1
*/
func (s *TagService) GetTagV1(GetTagV1QueryParams *GetTagV1QueryParams) (*ResponseTagGetTagV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/tag"

	queryString, _ := query.Values(GetTagV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseTagGetTagV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTagV1(GetTagV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetTagV1")
	}

	result := response.Result().(*ResponseTagGetTagV1)
	return result, response, err

}

//GetTagCountV1 Get Tag Count - 8091-a9b8-4bfb-a53b
/* Returns tag count


@param GetTagCountV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-tag-count-v1
*/
func (s *TagService) GetTagCountV1(GetTagCountV1QueryParams *GetTagCountV1QueryParams) (*ResponseTagGetTagCountV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/tag/count"

	queryString, _ := query.Values(GetTagCountV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseTagGetTagCountV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTagCountV1(GetTagCountV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetTagCountV1")
	}

	result := response.Result().(*ResponseTagGetTagCountV1)
	return result, response, err

}

//GetTagResourceTypesV1 Get Tag resource types - 4695-090d-403b-8eaa
/* Returns list of supported resource types



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-tag-resource-types-v1
*/
func (s *TagService) GetTagResourceTypesV1() (*ResponseTagGetTagResourceTypesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/tag/member/type"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseTagGetTagResourceTypesV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTagResourceTypesV1()
		}
		return nil, response, fmt.Errorf("error with operation GetTagResourceTypesV1")
	}

	result := response.Result().(*ResponseTagGetTagResourceTypesV1)
	return result, response, err

}

//GetTagByIDV1 Get Tag by Id - c1a3-59b1-4c89-b573
/* Returns tag specified by Id


@param id id path parameter. Tag ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-tag-by-id-v1
*/
func (s *TagService) GetTagByIDV1(id string) (*ResponseTagGetTagByIDV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/tag/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseTagGetTagByIDV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTagByIDV1(id)
		}
		return nil, response, fmt.Errorf("error with operation GetTagByIdV1")
	}

	result := response.Result().(*ResponseTagGetTagByIDV1)
	return result, response, err

}

//GetTagMembersByIDV1 Get Tag members by Id - eab7-abe0-48fb-99ad
/* Returns tag members specified by id


@param id id path parameter. Tag ID

@param GetTagMembersByIdV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-tag-members-by-id-v1
*/
func (s *TagService) GetTagMembersByIDV1(id string, GetTagMembersByIdV1QueryParams *GetTagMembersByIDV1QueryParams) (*ResponseTagGetTagMembersByIDV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/tag/{id}/member"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(GetTagMembersByIdV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseTagGetTagMembersByIDV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTagMembersByIDV1(id, GetTagMembersByIdV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetTagMembersByIdV1")
	}

	result := response.Result().(*ResponseTagGetTagMembersByIDV1)
	return result, response, err

}

//GetTagMemberCountV1 Get Tag Member count - 2e9d-b858-40fb-b1cf
/* Returns the number of members in a given tag


@param id id path parameter. Tag ID

@param GetTagMemberCountV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-tag-member-count-v1
*/
func (s *TagService) GetTagMemberCountV1(id string, GetTagMemberCountV1QueryParams *GetTagMemberCountV1QueryParams) (*ResponseTagGetTagMemberCountV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/tag/{id}/member/count"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(GetTagMemberCountV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseTagGetTagMemberCountV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTagMemberCountV1(id, GetTagMemberCountV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetTagMemberCountV1")
	}

	result := response.Result().(*ResponseTagGetTagMemberCountV1)
	return result, response, err

}

//RetrieveTagsAssociatedWithTheInterfacesV1 Retrieve tags associated with the interfaces. - b786-6abb-47f8-8c83
/* Fetches the tags associated with the interfaces. Interfaces that don't have any tags associated will not be included in the response. A tag is a user-defined or system-defined construct to group resources. When an interface is tagged, it is called a member of the tag.


@param RetrieveTagsAssociatedWithTheInterfacesV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieve-tags-associated-with-the-interfaces-v1
*/
func (s *TagService) RetrieveTagsAssociatedWithTheInterfacesV1(RetrieveTagsAssociatedWithTheInterfacesV1QueryParams *RetrieveTagsAssociatedWithTheInterfacesV1QueryParams) (*ResponseTagRetrieveTagsAssociatedWithTheInterfacesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/tags/interfaces/membersAssociations"

	queryString, _ := query.Values(RetrieveTagsAssociatedWithTheInterfacesV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseTagRetrieveTagsAssociatedWithTheInterfacesV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrieveTagsAssociatedWithTheInterfacesV1(RetrieveTagsAssociatedWithTheInterfacesV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrieveTagsAssociatedWithTheInterfacesV1")
	}

	result := response.Result().(*ResponseTagRetrieveTagsAssociatedWithTheInterfacesV1)
	return result, response, err

}

//RetrieveTheCountOfInterfacesThatAreAssociatedWithAtLeastOneTagV1 Retrieve the count of interfaces that are associated with at least one tag. - 3cb8-f8e6-4bfa-928c
/* Fetches the count of interfaces that are associated with at least one tag. A tag is a user-defined or system-defined construct to group resources. When an interface is tagged, it is called a member of the tag.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieve-the-count-of-interfaces-that-are-associated-with-at-least-one-tag-v1
*/
func (s *TagService) RetrieveTheCountOfInterfacesThatAreAssociatedWithAtLeastOneTagV1() (*ResponseTagRetrieveTheCountOfInterfacesThatAreAssociatedWithAtLeastOneTagV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/tags/interfaces/membersAssociations/count"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseTagRetrieveTheCountOfInterfacesThatAreAssociatedWithAtLeastOneTagV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrieveTheCountOfInterfacesThatAreAssociatedWithAtLeastOneTagV1()
		}
		return nil, response, fmt.Errorf("error with operation RetrieveTheCountOfInterfacesThatAreAssociatedWithAtLeastOneTagV1")
	}

	result := response.Result().(*ResponseTagRetrieveTheCountOfInterfacesThatAreAssociatedWithAtLeastOneTagV1)
	return result, response, err

}

//RetrieveTagsAssociatedWithNetworkDevicesV1 Retrieve tags associated with network devices. - 0b84-9b56-4bda-bc68
/* Fetches the tags associated with network devices. Devices that don't have any tags associated will not be included in the response. A tag is a user-defined or system-defined construct to group resources. When a device is tagged, it is called a member of the tag.


@param RetrieveTagsAssociatedWithNetworkDevicesV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieve-tags-associated-with-network-devices-v1
*/
func (s *TagService) RetrieveTagsAssociatedWithNetworkDevicesV1(RetrieveTagsAssociatedWithNetworkDevicesV1QueryParams *RetrieveTagsAssociatedWithNetworkDevicesV1QueryParams) (*ResponseTagRetrieveTagsAssociatedWithNetworkDevicesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/tags/networkDevices/membersAssociations"

	queryString, _ := query.Values(RetrieveTagsAssociatedWithNetworkDevicesV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseTagRetrieveTagsAssociatedWithNetworkDevicesV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrieveTagsAssociatedWithNetworkDevicesV1(RetrieveTagsAssociatedWithNetworkDevicesV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrieveTagsAssociatedWithNetworkDevicesV1")
	}

	result := response.Result().(*ResponseTagRetrieveTagsAssociatedWithNetworkDevicesV1)
	return result, response, err

}

//RetrieveTheCountOfNetworkDevicesThatAreAssociatedWithAtLeastOneTagV1 Retrieve the count of network devices that are associated with at least one tag. - 63b1-69be-4919-9dc0
/* Fetches the count of network devices that are associated with at least one tag. A tag is a user-defined or system-defined construct to group resources. When a device is tagged, it is called a member of the tag.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieve-the-count-of-network-devices-that-are-associated-with-at-least-one-tag-v1
*/
func (s *TagService) RetrieveTheCountOfNetworkDevicesThatAreAssociatedWithAtLeastOneTagV1() (*ResponseTagRetrieveTheCountOfNetworkDevicesThatAreAssociatedWithAtLeastOneTagV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/tags/networkDevices/membersAssociations/count"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseTagRetrieveTheCountOfNetworkDevicesThatAreAssociatedWithAtLeastOneTagV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrieveTheCountOfNetworkDevicesThatAreAssociatedWithAtLeastOneTagV1()
		}
		return nil, response, fmt.Errorf("error with operation RetrieveTheCountOfNetworkDevicesThatAreAssociatedWithAtLeastOneTagV1")
	}

	result := response.Result().(*ResponseTagRetrieveTheCountOfNetworkDevicesThatAreAssociatedWithAtLeastOneTagV1)
	return result, response, err

}

//CreateTagV1 Create Tag - 1399-891c-42a8-be64
/* Creates tag with specified tag attributes



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-tag-v1
*/
func (s *TagService) CreateTagV1(requestTagCreateTagV1 *RequestTagCreateTagV1) (*ResponseTagCreateTagV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/tag"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestTagCreateTagV1).
		SetResult(&ResponseTagCreateTagV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateTagV1(requestTagCreateTagV1)
		}

		return nil, response, fmt.Errorf("error with operation CreateTagV1")
	}

	result := response.Result().(*ResponseTagCreateTagV1)
	return result, response, err

}

//AddMembersToTheTagV1 Add members to the tag - 00a2-fa61-4608-9317
/* Adds members to the tag specified by id


@param id id path parameter. Tag ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-members-to-the-tag-v1
*/
func (s *TagService) AddMembersToTheTagV1(id string, requestTagAddMembersToTheTagV1 *RequestTagAddMembersToTheTagV1) (*ResponseTagAddMembersToTheTagV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/tag/{id}/member"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestTagAddMembersToTheTagV1).
		SetResult(&ResponseTagAddMembersToTheTagV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddMembersToTheTagV1(id, requestTagAddMembersToTheTagV1)
		}

		return nil, response, fmt.Errorf("error with operation AddMembersToTheTagV1")
	}

	result := response.Result().(*ResponseTagAddMembersToTheTagV1)
	return result, response, err

}

//QueryTheTagsAssociatedWithInterfacesV1 Query the tags associated with interfaces. - 87a2-4a4e-4109-becf
/* Fetches the tags associated with the given interface `ids`. Interfaces that don't have any tags associated will not be included in the response. A tag is a user-defined or system-defined construct to group resources. When an interface is tagged, it is called a member of the tag. `ids` can be fetched via `/dna/intent/api/v1/interface` API.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!query-the-tags-associated-with-interfaces-v1
*/
func (s *TagService) QueryTheTagsAssociatedWithInterfacesV1(requestTagQueryTheTagsAssociatedWithInterfacesV1 *RequestTagQueryTheTagsAssociatedWithInterfacesV1) (*ResponseTagQueryTheTagsAssociatedWithInterfacesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/tags/interfaces/membersAssociations/query"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestTagQueryTheTagsAssociatedWithInterfacesV1).
		SetResult(&ResponseTagQueryTheTagsAssociatedWithInterfacesV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.QueryTheTagsAssociatedWithInterfacesV1(requestTagQueryTheTagsAssociatedWithInterfacesV1)
		}

		return nil, response, fmt.Errorf("error with operation QueryTheTagsAssociatedWithInterfacesV1")
	}

	result := response.Result().(*ResponseTagQueryTheTagsAssociatedWithInterfacesV1)
	return result, response, err

}

//QueryTheTagsAssociatedWithNetworkDevicesV1 Query the tags associated with network devices. - 6480-fa01-417b-b397
/* Fetches the tags associated with the given network device `ids`. Devices that don't have any tags associated will not be included in the response. A tag is a user-defined or system-defined construct to group resources. When a device is tagged, it is called a member of the tag. `ids` can be fetched via `/dna/intent/api/v1/network-device` API.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!query-the-tags-associated-with-network-devices-v1
*/
func (s *TagService) QueryTheTagsAssociatedWithNetworkDevicesV1(requestTagQueryTheTagsAssociatedWithNetworkDevicesV1 *RequestTagQueryTheTagsAssociatedWithNetworkDevicesV1) (*ResponseTagQueryTheTagsAssociatedWithNetworkDevicesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/tags/networkDevices/membersAssociations/query"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestTagQueryTheTagsAssociatedWithNetworkDevicesV1).
		SetResult(&ResponseTagQueryTheTagsAssociatedWithNetworkDevicesV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.QueryTheTagsAssociatedWithNetworkDevicesV1(requestTagQueryTheTagsAssociatedWithNetworkDevicesV1)
		}

		return nil, response, fmt.Errorf("error with operation QueryTheTagsAssociatedWithNetworkDevicesV1")
	}

	result := response.Result().(*ResponseTagQueryTheTagsAssociatedWithNetworkDevicesV1)
	return result, response, err

}

//UpdateTagV1 Update Tag - 4d86-a993-469a-9da9
/* Updates a tag specified by id


 */
func (s *TagService) UpdateTagV1(requestTagUpdateTagV1 *RequestTagUpdateTagV1) (*ResponseTagUpdateTagV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/tag"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestTagUpdateTagV1).
		SetResult(&ResponseTagUpdateTagV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateTagV1(requestTagUpdateTagV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdateTagV1")
	}

	result := response.Result().(*ResponseTagUpdateTagV1)
	return result, response, err

}

//UpdateTagMembershipV1 Update tag membership - 45bc-7a83-44a8-bc1e
/* Update tag membership. As part of the request payload through this API, only the specified members are added / retained to the given input tags. Possible values of memberType attribute in the request payload can be queried by using the /tag/member/type API


 */
func (s *TagService) UpdateTagMembershipV1(requestTagUpdateTagMembershipV1 *RequestTagUpdateTagMembershipV1) (*ResponseTagUpdateTagMembershipV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/tag/member"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestTagUpdateTagMembershipV1).
		SetResult(&ResponseTagUpdateTagMembershipV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateTagMembershipV1(requestTagUpdateTagMembershipV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdateTagMembershipV1")
	}

	result := response.Result().(*ResponseTagUpdateTagMembershipV1)
	return result, response, err

}

//DeleteTagV1 Delete Tag - 429c-2815-4bda-a13d
/* Deletes a tag specified by id


@param id id path parameter. Tag ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-tag-v1
*/
func (s *TagService) DeleteTagV1(id string) (*ResponseTagDeleteTagV1, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/tag/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseTagDeleteTagV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteTagV1(id)
		}
		return nil, response, fmt.Errorf("error with operation DeleteTagV1")
	}

	result := response.Result().(*ResponseTagDeleteTagV1)
	return result, response, err

}

//RemoveTagMemberV1 Remove Tag member - caa3-ea70-4d78-b37e
/* Removes Tag member from the tag specified by id


@param id id path parameter. Tag ID

@param memberID memberId path parameter. TagMember id to be removed from tag


Documentation Link: https://developer.cisco.com/docs/dna-center/#!remove-tag-member-v1
*/
func (s *TagService) RemoveTagMemberV1(id string, memberID string) (*ResponseTagRemoveTagMemberV1, *resty.Response, error) {
	//id string,memberID string
	path := "/dna/intent/api/v1/tag/{id}/member/{memberId}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)
	path = strings.Replace(path, "{memberId}", fmt.Sprintf("%v", memberID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseTagRemoveTagMemberV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RemoveTagMemberV1(id, memberID)
		}
		return nil, response, fmt.Errorf("error with operation RemoveTagMemberV1")
	}

	result := response.Result().(*ResponseTagRemoveTagMemberV1)
	return result, response, err

}

// Alias Function
/*
This method acts as an alias for the method `GetTagMembersByIDV1`
*/
func (s *TagService) GetTagMembersByID(id string, GetTagMembersByIdV1QueryParams *GetTagMembersByIDV1QueryParams) (*ResponseTagGetTagMembersByIDV1, *resty.Response, error) {
	return s.GetTagMembersByIDV1(id, GetTagMembersByIdV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `UpdateTagV1`
*/
func (s *TagService) UpdateTag(requestTagUpdateTagV1 *RequestTagUpdateTagV1) (*ResponseTagUpdateTagV1, *resty.Response, error) {
	return s.UpdateTagV1(requestTagUpdateTagV1)
}

// Alias Function
/*
This method acts as an alias for the method `DeleteTagV1`
*/
func (s *TagService) DeleteTag(id string) (*ResponseTagDeleteTagV1, *resty.Response, error) {
	return s.DeleteTagV1(id)
}

// Alias Function
/*
This method acts as an alias for the method `GetTagV1`
*/
func (s *TagService) GetTag(GetTagV1QueryParams *GetTagV1QueryParams) (*ResponseTagGetTagV1, *resty.Response, error) {
	return s.GetTagV1(GetTagV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `RemoveTagMemberV1`
*/
func (s *TagService) RemoveTagMember(id string, memberID string) (*ResponseTagRemoveTagMemberV1, *resty.Response, error) {
	return s.RemoveTagMemberV1(id, memberID)
}

// Alias Function
/*
This method acts as an alias for the method `RetrieveTagsAssociatedWithTheInterfacesV1`
*/
func (s *TagService) RetrieveTagsAssociatedWithTheInterfaces(RetrieveTagsAssociatedWithTheInterfacesV1QueryParams *RetrieveTagsAssociatedWithTheInterfacesV1QueryParams) (*ResponseTagRetrieveTagsAssociatedWithTheInterfacesV1, *resty.Response, error) {
	return s.RetrieveTagsAssociatedWithTheInterfacesV1(RetrieveTagsAssociatedWithTheInterfacesV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetTagResourceTypesV1`
*/
func (s *TagService) GetTagResourceTypes() (*ResponseTagGetTagResourceTypesV1, *resty.Response, error) {
	return s.GetTagResourceTypesV1()
}

// Alias Function
/*
This method acts as an alias for the method `RetrieveTheCountOfNetworkDevicesThatAreAssociatedWithAtLeastOneTagV1`
*/
func (s *TagService) RetrieveTheCountOfNetworkDevicesThatAreAssociatedWithAtLeastOneTag() (*ResponseTagRetrieveTheCountOfNetworkDevicesThatAreAssociatedWithAtLeastOneTagV1, *resty.Response, error) {
	return s.RetrieveTheCountOfNetworkDevicesThatAreAssociatedWithAtLeastOneTagV1()
}

// Alias Function
/*
This method acts as an alias for the method `QueryTheTagsAssociatedWithNetworkDevicesV1`
*/
func (s *TagService) QueryTheTagsAssociatedWithNetworkDevices(requestTagQueryTheTagsAssociatedWithNetworkDevicesV1 *RequestTagQueryTheTagsAssociatedWithNetworkDevicesV1) (*ResponseTagQueryTheTagsAssociatedWithNetworkDevicesV1, *resty.Response, error) {
	return s.QueryTheTagsAssociatedWithNetworkDevicesV1(requestTagQueryTheTagsAssociatedWithNetworkDevicesV1)
}

// Alias Function
/*
This method acts as an alias for the method `RetrieveTheCountOfInterfacesThatAreAssociatedWithAtLeastOneTagV1`
*/
func (s *TagService) RetrieveTheCountOfInterfacesThatAreAssociatedWithAtLeastOneTag() (*ResponseTagRetrieveTheCountOfInterfacesThatAreAssociatedWithAtLeastOneTagV1, *resty.Response, error) {
	return s.RetrieveTheCountOfInterfacesThatAreAssociatedWithAtLeastOneTagV1()
}

// Alias Function
/*
This method acts as an alias for the method `CreateTagV1`
*/
func (s *TagService) CreateTag(requestTagCreateTagV1 *RequestTagCreateTagV1) (*ResponseTagCreateTagV1, *resty.Response, error) {
	return s.CreateTagV1(requestTagCreateTagV1)
}

// Alias Function
/*
This method acts as an alias for the method `RetrieveTagsAssociatedWithNetworkDevicesV1`
*/
func (s *TagService) RetrieveTagsAssociatedWithNetworkDevices(RetrieveTagsAssociatedWithNetworkDevicesV1QueryParams *RetrieveTagsAssociatedWithNetworkDevicesV1QueryParams) (*ResponseTagRetrieveTagsAssociatedWithNetworkDevicesV1, *resty.Response, error) {
	return s.RetrieveTagsAssociatedWithNetworkDevicesV1(RetrieveTagsAssociatedWithNetworkDevicesV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `GetTagMemberCountV1`
*/
func (s *TagService) GetTagMemberCount(id string, GetTagMemberCountV1QueryParams *GetTagMemberCountV1QueryParams) (*ResponseTagGetTagMemberCountV1, *resty.Response, error) {
	return s.GetTagMemberCountV1(id, GetTagMemberCountV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `QueryTheTagsAssociatedWithInterfacesV1`
*/
func (s *TagService) QueryTheTagsAssociatedWithInterfaces(requestTagQueryTheTagsAssociatedWithInterfacesV1 *RequestTagQueryTheTagsAssociatedWithInterfacesV1) (*ResponseTagQueryTheTagsAssociatedWithInterfacesV1, *resty.Response, error) {
	return s.QueryTheTagsAssociatedWithInterfacesV1(requestTagQueryTheTagsAssociatedWithInterfacesV1)
}

// Alias Function
/*
This method acts as an alias for the method `UpdateTagMembershipV1`
*/
func (s *TagService) UpdateTagMembership(requestTagUpdateTagMembershipV1 *RequestTagUpdateTagMembershipV1) (*ResponseTagUpdateTagMembershipV1, *resty.Response, error) {
	return s.UpdateTagMembershipV1(requestTagUpdateTagMembershipV1)
}

// Alias Function
/*
This method acts as an alias for the method `GetTagCountV1`
*/
func (s *TagService) GetTagCount(GetTagCountV1QueryParams *GetTagCountV1QueryParams) (*ResponseTagGetTagCountV1, *resty.Response, error) {
	return s.GetTagCountV1(GetTagCountV1QueryParams)
}

// Alias Function
/*
This method acts as an alias for the method `AddMembersToTheTagV1`
*/
func (s *TagService) AddMembersToTheTag(id string, requestTagAddMembersToTheTagV1 *RequestTagAddMembersToTheTagV1) (*ResponseTagAddMembersToTheTagV1, *resty.Response, error) {
	return s.AddMembersToTheTagV1(id, requestTagAddMembersToTheTagV1)
}

// Alias Function
/*
This method acts as an alias for the method `GetTagByIDV1`
*/
func (s *TagService) GetTagByID(id string) (*ResponseTagGetTagByIDV1, *resty.Response, error) {
	return s.GetTagByIDV1(id)
}
