package catalyst

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type ConfigurationTemplatesService service

type CreatesACloneOfTheGivenTemplateV1QueryParams struct {
	ProjectID string `url:"projectId,omitempty"` //UUID of the project in which the template needs to be created
}
type GetsAListOfProjectsV1QueryParams struct {
	Name      string `url:"name,omitempty"`      //Name of project to be searched
	SortOrder string `url:"sortOrder,omitempty"` //Sort Order Ascending (asc) or Descending (des)
}
type ImportsTheProjectsProvidedV1QueryParams struct {
	DoVersion bool `url:"doVersion,omitempty"` //If this flag is true then it creates a new version of the template with the imported contents in case if the templates already exists. " If this flag is false and if template already exists, then operation fails with 'Template already exists' error
}
type ImportsTheTemplatesProvidedV1QueryParams struct {
	DoVersion bool `url:"doVersion,omitempty"` //If this flag is true then it creates a new version of the template with the imported contents in case if the templates already exists. " If this flag is false and if template already exists, then operation fails with 'Template already exists' error
}
type GetsTheTemplatesAvailableV1QueryParams struct {
	ProjectID                  string   `url:"projectId,omitempty"`                  //Filter template(s) based on project UUID
	SoftwareType               string   `url:"softwareType,omitempty"`               //Filter template(s) based software type
	SoftwareVersion            string   `url:"softwareVersion,omitempty"`            //Filter template(s) based softwareVersion
	ProductFamily              string   `url:"productFamily,omitempty"`              //Filter template(s) based on device family
	ProductSeries              string   `url:"productSeries,omitempty"`              //Filter template(s) based on device series
	ProductType                string   `url:"productType,omitempty"`                //Filter template(s) based on device type
	FilterConflictingTemplates bool     `url:"filterConflictingTemplates,omitempty"` //Filter template(s) based on confliting templates
	Tags                       []string `url:"tags,omitempty"`                       //Filter template(s) based on tags
	ProjectNames               []string `url:"projectNames,omitempty"`               //Filter template(s) based on project names
	UnCommitted                bool     `url:"unCommitted,omitempty"`                //Filter template(s) based on template commited or not
	SortOrder                  string   `url:"sortOrder,omitempty"`                  //Sort Order Ascending (asc) or Descending (des)
}
type GetsDetailsOfAGivenTemplateV1QueryParams struct {
	LatestVersion bool `url:"latestVersion,omitempty"` //latestVersion flag to get the latest versioned template
}
type GetProjectsDetailsV2QueryParams struct {
	ID        string `url:"id,omitempty"`        //Id of project to be searched
	Name      string `url:"name,omitempty"`      //Name of project to be searched
	Offset    int    `url:"offset,omitempty"`    //Index of first result
	Limit     int    `url:"limit,omitempty"`     //Limits number of results
	SortOrder string `url:"sortOrder,omitempty"` //Sort Order Ascending (asc) or Descending (dsc)
}
type GetTemplatesDetailsV2QueryParams struct {
	ID                         string   `url:"id,omitempty"`                         //Id of template to be searched
	Name                       string   `url:"name,omitempty"`                       //Name of template to be searched
	ProjectID                  string   `url:"projectId,omitempty"`                  //Filter template(s) based on project id
	ProjectName                string   `url:"projectName,omitempty"`                //Filter template(s) based on project name
	SoftwareType               string   `url:"softwareType,omitempty"`               //Filter template(s) based software type
	SoftwareVersion            string   `url:"softwareVersion,omitempty"`            //Filter template(s) based softwareVersion
	ProductFamily              string   `url:"productFamily,omitempty"`              //Filter template(s) based on device family
	ProductSeries              string   `url:"productSeries,omitempty"`              //Filter template(s) based on device series
	ProductType                string   `url:"productType,omitempty"`                //Filter template(s) based on device type
	FilterConflictingTemplates bool     `url:"filterConflictingTemplates,omitempty"` //Filter template(s) based on confliting templates
	Tags                       []string `url:"tags,omitempty"`                       //Filter template(s) based on tags
	UnCommitted                bool     `url:"unCommitted,omitempty"`                //Return uncommitted template
	SortOrder                  string   `url:"sortOrder,omitempty"`                  //Sort Order Ascending (asc) or Descending (dsc)
	AllTemplateAttributes      bool     `url:"allTemplateAttributes,omitempty"`      //Return all template attributes
	IncludeVersionDetails      bool     `url:"includeVersionDetails,omitempty"`      //Include template version details
	Offset                     int      `url:"offset,omitempty"`                     //Index of first result
	Limit                      int      `url:"limit,omitempty"`                      //Limits number of results
}

type ResponseConfigurationTemplatesCreatesACloneOfTheGivenTemplateV1 struct {
	Response *ResponseConfigurationTemplatesCreatesACloneOfTheGivenTemplateV1Response `json:"response,omitempty"` //
	Version  string                                                                   `json:"version,omitempty"`  //
}
type ResponseConfigurationTemplatesCreatesACloneOfTheGivenTemplateV1Response struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseConfigurationTemplatesCreateProjectV1 struct {
	Response *ResponseConfigurationTemplatesCreateProjectV1Response `json:"response,omitempty"` //
	Version  string                                                 `json:"version,omitempty"`  //
}
type ResponseConfigurationTemplatesCreateProjectV1Response struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseConfigurationTemplatesUpdateProjectV1 struct {
	Response *ResponseConfigurationTemplatesUpdateProjectV1Response `json:"response,omitempty"` //
	Version  string                                                 `json:"version,omitempty"`  //
}
type ResponseConfigurationTemplatesUpdateProjectV1Response struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseConfigurationTemplatesGetsAListOfProjectsV1 []ResponseItemConfigurationTemplatesGetsAListOfProjectsV1 // Array of ResponseConfigurationTemplatesGetsAListOfProjectsV1
type ResponseItemConfigurationTemplatesGetsAListOfProjectsV1 struct {
	Tags           *[]ResponseItemConfigurationTemplatesGetsAListOfProjectsV1Tags    `json:"tags,omitempty"`           //
	CreateTime     *int                                                              `json:"createTime,omitempty"`     // Create time of project
	Description    string                                                            `json:"description,omitempty"`    // Description of project
	ID             string                                                            `json:"id,omitempty"`             // UUID of project
	LastUpdateTime *int                                                              `json:"lastUpdateTime,omitempty"` // Update time of project
	Name           string                                                            `json:"name,omitempty"`           // Name of project
	Templates      *ResponseItemConfigurationTemplatesGetsAListOfProjectsV1Templates `json:"templates,omitempty"`      // List of templates within the project
}
type ResponseItemConfigurationTemplatesGetsAListOfProjectsV1Tags struct {
	ID   string `json:"id,omitempty"`   // UUID of tag
	Name string `json:"name,omitempty"` // Name of tag
}
type ResponseItemConfigurationTemplatesGetsAListOfProjectsV1Templates struct {
	Tags                    *[]ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesTags                   `json:"tags,omitempty"`                    //
	Author                  string                                                                                  `json:"author,omitempty"`                  // Author of template
	Composite               *bool                                                                                   `json:"composite,omitempty"`               // Is it composite template
	ContainingTemplates     *[]ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesContainingTemplates    `json:"containingTemplates,omitempty"`     //
	CreateTime              *int                                                                                    `json:"createTime,omitempty"`              // Create time of template
	CustomParamsOrder       *bool                                                                                   `json:"customParamsOrder,omitempty"`       // Custom Params Order
	Description             string                                                                                  `json:"description,omitempty"`             // Description of template
	DeviceTypes             *[]ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesDeviceTypes            `json:"deviceTypes,omitempty"`             //
	FailurePolicy           string                                                                                  `json:"failurePolicy,omitempty"`           // Define failure policy if template provisioning fails
	ID                      string                                                                                  `json:"id,omitempty"`                      // UUID of template
	Language                string                                                                                  `json:"language,omitempty"`                // Template language (JINJA or VELOCITY)
	LastUpdateTime          *int                                                                                    `json:"lastUpdateTime,omitempty"`          // Update time of template
	LatestVersionTime       *int                                                                                    `json:"latestVersionTime,omitempty"`       // Latest versioned template time
	Name                    string                                                                                  `json:"name,omitempty"`                    // Name of template
	ParentTemplateID        string                                                                                  `json:"parentTemplateId,omitempty"`        // Parent templateID
	ProjectID               string                                                                                  `json:"projectId,omitempty"`               // Project UUID
	ProjectName             string                                                                                  `json:"projectName,omitempty"`             // Project name
	RollbackTemplateContent string                                                                                  `json:"rollbackTemplateContent,omitempty"` // Rollback template content
	RollbackTemplateParams  *[]ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesRollbackTemplateParams `json:"rollbackTemplateParams,omitempty"`  //
	SoftwareType            string                                                                                  `json:"softwareType,omitempty"`            // Applicable device software type
	SoftwareVariant         string                                                                                  `json:"softwareVariant,omitempty"`         // Applicable device software variant
	SoftwareVersion         string                                                                                  `json:"softwareVersion,omitempty"`         // Applicable device software version
	TemplateContent         string                                                                                  `json:"templateContent,omitempty"`         // Template content
	TemplateParams          *[]ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesTemplateParams         `json:"templateParams,omitempty"`          //
	ValidationErrors        *ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesValidationErrors         `json:"validationErrors,omitempty"`        //
	Version                 string                                                                                  `json:"version,omitempty"`                 // Current version of template
}
type ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesTags struct {
	ID   string `json:"id,omitempty"`   // UUID of tag
	Name string `json:"name,omitempty"` // Name of tag
}
type ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesContainingTemplates struct {
	Tags                   *[]ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesContainingTemplatesTags                   `json:"tags,omitempty"`                   //
	Composite              *bool                                                                                                      `json:"composite,omitempty"`              // Is it composite template
	Description            string                                                                                                     `json:"description,omitempty"`            // Description of template
	DeviceTypes            *[]ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesContainingTemplatesDeviceTypes            `json:"deviceTypes,omitempty"`            //
	ID                     string                                                                                                     `json:"id,omitempty"`                     // UUID of template
	Language               string                                                                                                     `json:"language,omitempty"`               // Template language (JINJA or VELOCITY)
	Name                   string                                                                                                     `json:"name,omitempty"`                   // Name of template
	ProjectName            string                                                                                                     `json:"projectName,omitempty"`            // Project name
	RollbackTemplateParams *[]ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesContainingTemplatesRollbackTemplateParams `json:"rollbackTemplateParams,omitempty"` //
	TemplateContent        string                                                                                                     `json:"templateContent,omitempty"`        // Template content
	TemplateParams         *[]ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesContainingTemplatesTemplateParams         `json:"templateParams,omitempty"`         //
	Version                string                                                                                                     `json:"version,omitempty"`                // Current version of template
}
type ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesContainingTemplatesTags struct {
	ID   string `json:"id,omitempty"`   // UUID of tag
	Name string `json:"name,omitempty"` // Name of tag
}
type ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesContainingTemplatesDeviceTypes struct {
	ProductFamily string `json:"productFamily,omitempty"` // Device family
	ProductSeries string `json:"productSeries,omitempty"` // Device series
	ProductType   string `json:"productType,omitempty"`   // Device type
}
type ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesContainingTemplatesRollbackTemplateParams struct {
	Binding         string                                                                                                            `json:"binding,omitempty"`         // Bind to source
	CustomOrder     *int                                                                                                              `json:"customOrder,omitempty"`     // CustomOrder of template param
	DataType        string                                                                                                            `json:"dataType,omitempty"`        // Datatype of template param
	DefaultValue    string                                                                                                            `json:"defaultValue,omitempty"`    // Default value of template param
	Description     string                                                                                                            `json:"description,omitempty"`     // Description of template param
	DisplayName     string                                                                                                            `json:"displayName,omitempty"`     // Display name of param
	Group           string                                                                                                            `json:"group,omitempty"`           // group
	ID              string                                                                                                            `json:"id,omitempty"`              // UUID of template param
	InstructionText string                                                                                                            `json:"instructionText,omitempty"` // Instruction text for param
	Key             string                                                                                                            `json:"key,omitempty"`             // key
	NotParam        *bool                                                                                                             `json:"notParam,omitempty"`        // Is it not a variable
	Order           *int                                                                                                              `json:"order,omitempty"`           // Order of template param
	ParamArray      *bool                                                                                                             `json:"paramArray,omitempty"`      // Is it an array
	ParameterName   string                                                                                                            `json:"parameterName,omitempty"`   // Name of template param
	Provider        string                                                                                                            `json:"provider,omitempty"`        // provider
	Range           *[]ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesContainingTemplatesRollbackTemplateParamsRange   `json:"range,omitempty"`           //
	Required        *bool                                                                                                             `json:"required,omitempty"`        // Is param required
	Selection       *ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesContainingTemplatesRollbackTemplateParamsSelection `json:"selection,omitempty"`       //
}
type ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesContainingTemplatesRollbackTemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       // UUID of range
	MaxValue *int   `json:"maxValue,omitempty"` // Max value of range
	MinValue *int   `json:"minValue,omitempty"` // Min value of range
}
type ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesContainingTemplatesRollbackTemplateParamsSelection struct {
	DefaultSelectedValues []string                                                                                                                         `json:"defaultSelectedValues,omitempty"` // Default selection values
	ID                    string                                                                                                                           `json:"id,omitempty"`                    // UUID of selection
	SelectionType         string                                                                                                                           `json:"selectionType,omitempty"`         // Type of selection(SINGLE_SELECT or MULTI_SELECT)
	SelectionValues       *ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesContainingTemplatesRollbackTemplateParamsSelectionSelectionValues `json:"selectionValues,omitempty"`       // Selection values
}
type ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesContainingTemplatesRollbackTemplateParamsSelectionSelectionValues interface{}
type ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesContainingTemplatesTemplateParams struct {
	Binding         string                                                                                                    `json:"binding,omitempty"`         // Bind to source
	CustomOrder     *int                                                                                                      `json:"customOrder,omitempty"`     // CustomOrder of template param
	DataType        string                                                                                                    `json:"dataType,omitempty"`        // Datatype of template param
	DefaultValue    string                                                                                                    `json:"defaultValue,omitempty"`    // Default value of template param
	Description     string                                                                                                    `json:"description,omitempty"`     // Description of template param
	DisplayName     string                                                                                                    `json:"displayName,omitempty"`     // Display name of param
	Group           string                                                                                                    `json:"group,omitempty"`           // group
	ID              string                                                                                                    `json:"id,omitempty"`              // UUID of template param
	InstructionText string                                                                                                    `json:"instructionText,omitempty"` // Instruction text for param
	Key             string                                                                                                    `json:"key,omitempty"`             // key
	NotParam        *bool                                                                                                     `json:"notParam,omitempty"`        // Is it not a variable
	Order           *int                                                                                                      `json:"order,omitempty"`           // Order of template param
	ParamArray      *bool                                                                                                     `json:"paramArray,omitempty"`      // Is it an array
	ParameterName   string                                                                                                    `json:"parameterName,omitempty"`   // Name of template param
	Provider        string                                                                                                    `json:"provider,omitempty"`        // provider
	Range           *[]ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesContainingTemplatesTemplateParamsRange   `json:"range,omitempty"`           //
	Required        *bool                                                                                                     `json:"required,omitempty"`        // Is param required
	Selection       *ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesContainingTemplatesTemplateParamsSelection `json:"selection,omitempty"`       //
}
type ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesContainingTemplatesTemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       // UUID of range
	MaxValue *int   `json:"maxValue,omitempty"` // Max value of range
	MinValue *int   `json:"minValue,omitempty"` // Min value of range
}
type ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesContainingTemplatesTemplateParamsSelection struct {
	DefaultSelectedValues []string                                                                                                                 `json:"defaultSelectedValues,omitempty"` // Default selection values
	ID                    string                                                                                                                   `json:"id,omitempty"`                    // UUID of selection
	SelectionType         string                                                                                                                   `json:"selectionType,omitempty"`         // Type of selection(SINGLE_SELECT or MULTI_SELECT)
	SelectionValues       *ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesContainingTemplatesTemplateParamsSelectionSelectionValues `json:"selectionValues,omitempty"`       // Selection values
}
type ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesContainingTemplatesTemplateParamsSelectionSelectionValues interface{}
type ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesDeviceTypes struct {
	ProductFamily string `json:"productFamily,omitempty"` // Device family
	ProductSeries string `json:"productSeries,omitempty"` // Device series
	ProductType   string `json:"productType,omitempty"`   // Device type
}
type ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesRollbackTemplateParams struct {
	Binding         string                                                                                         `json:"binding,omitempty"`         // Bind to source
	CustomOrder     *int                                                                                           `json:"customOrder,omitempty"`     // CustomOrder of template param
	DataType        string                                                                                         `json:"dataType,omitempty"`        // Datatype of template param
	DefaultValue    string                                                                                         `json:"defaultValue,omitempty"`    // Default value of template param
	Description     string                                                                                         `json:"description,omitempty"`     // Description of template param
	DisplayName     string                                                                                         `json:"displayName,omitempty"`     // Display name of param
	Group           string                                                                                         `json:"group,omitempty"`           // group
	ID              string                                                                                         `json:"id,omitempty"`              // UUID of template param
	InstructionText string                                                                                         `json:"instructionText,omitempty"` // Instruction text for param
	Key             string                                                                                         `json:"key,omitempty"`             // key
	NotParam        *bool                                                                                          `json:"notParam,omitempty"`        // Is it not a variable
	Order           *int                                                                                           `json:"order,omitempty"`           // Order of template param
	ParamArray      *bool                                                                                          `json:"paramArray,omitempty"`      // Is it an array
	ParameterName   string                                                                                         `json:"parameterName,omitempty"`   // Name of template param
	Provider        string                                                                                         `json:"provider,omitempty"`        // provider
	Range           *[]ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesRollbackTemplateParamsRange   `json:"range,omitempty"`           //
	Required        *bool                                                                                          `json:"required,omitempty"`        // Is param required
	Selection       *ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesRollbackTemplateParamsSelection `json:"selection,omitempty"`       //
}
type ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesRollbackTemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       // UUID of range
	MaxValue *int   `json:"maxValue,omitempty"` // Max value of range
	MinValue *int   `json:"minValue,omitempty"` // Min value of range
}
type ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesRollbackTemplateParamsSelection struct {
	DefaultSelectedValues []string                                                                                                      `json:"defaultSelectedValues,omitempty"` // Default selection values
	ID                    string                                                                                                        `json:"id,omitempty"`                    // UUID of selection
	SelectionType         string                                                                                                        `json:"selectionType,omitempty"`         // Type of selection(SINGLE_SELECT or MULTI_SELECT)
	SelectionValues       *ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesRollbackTemplateParamsSelectionSelectionValues `json:"selectionValues,omitempty"`       // Selection values
}
type ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesRollbackTemplateParamsSelectionSelectionValues interface{}
type ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesTemplateParams struct {
	Binding         string                                                                                 `json:"binding,omitempty"`         // Bind to source
	CustomOrder     *int                                                                                   `json:"customOrder,omitempty"`     // CustomOrder of template param
	DataType        string                                                                                 `json:"dataType,omitempty"`        // Datatype of template param
	DefaultValue    string                                                                                 `json:"defaultValue,omitempty"`    // Default value of template param
	Description     string                                                                                 `json:"description,omitempty"`     // Description of template param
	DisplayName     string                                                                                 `json:"displayName,omitempty"`     // Display name of param
	Group           string                                                                                 `json:"group,omitempty"`           // group
	ID              string                                                                                 `json:"id,omitempty"`              // UUID of template param
	InstructionText string                                                                                 `json:"instructionText,omitempty"` // Instruction text for param
	Key             string                                                                                 `json:"key,omitempty"`             // key
	NotParam        *bool                                                                                  `json:"notParam,omitempty"`        // Is it not a variable
	Order           *int                                                                                   `json:"order,omitempty"`           // Order of template param
	ParamArray      *bool                                                                                  `json:"paramArray,omitempty"`      // Is it an array
	ParameterName   string                                                                                 `json:"parameterName,omitempty"`   // Name of template param
	Provider        string                                                                                 `json:"provider,omitempty"`        // provider
	Range           *[]ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesTemplateParamsRange   `json:"range,omitempty"`           //
	Required        *bool                                                                                  `json:"required,omitempty"`        // Is param required
	Selection       *ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesTemplateParamsSelection `json:"selection,omitempty"`       //
}
type ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesTemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       // UUID of range
	MaxValue *int   `json:"maxValue,omitempty"` // Max value of range
	MinValue *int   `json:"minValue,omitempty"` // Min value of range
}
type ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesTemplateParamsSelection struct {
	DefaultSelectedValues []string                                                                                              `json:"defaultSelectedValues,omitempty"` // Default selection values
	ID                    string                                                                                                `json:"id,omitempty"`                    // UUID of selection
	SelectionType         string                                                                                                `json:"selectionType,omitempty"`         // Type of selection(SINGLE_SELECT or MULTI_SELECT)
	SelectionValues       *ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesTemplateParamsSelectionSelectionValues `json:"selectionValues,omitempty"`       // Selection values
}
type ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesTemplateParamsSelectionSelectionValues interface{}
type ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesValidationErrors struct {
	RollbackTemplateErrors *ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesValidationErrorsRollbackTemplateErrors `json:"rollbackTemplateErrors,omitempty"` // Validation or design conflicts errors of rollback template
	TemplateErrors         *ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesValidationErrorsTemplateErrors         `json:"templateErrors,omitempty"`         // Validation or design conflicts errors
	TemplateID             string                                                                                                `json:"templateId,omitempty"`             // UUID of template
	TemplateVersion        string                                                                                                `json:"templateVersion,omitempty"`        // Current version of template
}
type ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesValidationErrorsRollbackTemplateErrors interface{}
type ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesValidationErrorsTemplateErrors interface{}
type ResponseConfigurationTemplatesImportsTheProjectsProvidedV1 struct {
	Response *ResponseConfigurationTemplatesImportsTheProjectsProvidedResponse `json:"response,omitempty"` //
	Version  string                                                            `json:"version,omitempty"`  //
}
type ResponseConfigurationTemplatesImportsTheProjectsProvidedResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseConfigurationTemplatesExportsTheProjectsForAGivenCriteriaV1 struct {
	Response *ResponseConfigurationTemplatesExportsTheProjectsForAGivenCriteriaV1Response `json:"response,omitempty"` //
	Version  string                                                                       `json:"version,omitempty"`  //
}
type ResponseConfigurationTemplatesExportsTheProjectsForAGivenCriteriaV1Response struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseConfigurationTemplatesImportsTheTemplatesProvidedV1 struct {
	Response *ResponseConfigurationTemplatesImportsTheTemplatesProvidedV1Response `json:"response,omitempty"` //
	Version  string                                                               `json:"version,omitempty"`  //
}
type ResponseConfigurationTemplatesImportsTheTemplatesProvidedV1Response struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectV1 struct {
	Tags           *[]ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectV1Tags    `json:"tags,omitempty"`           //
	CreateTime     *int                                                                    `json:"createTime,omitempty"`     // Create time of project
	Description    string                                                                  `json:"description,omitempty"`    // Description of project
	ID             string                                                                  `json:"id,omitempty"`             // UUID of project
	LastUpdateTime *int                                                                    `json:"lastUpdateTime,omitempty"` // Update time of project
	Name           string                                                                  `json:"name,omitempty"`           // Name of project
	Templates      *ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectV1Templates `json:"templates,omitempty"`      // List of templates within the project
}
type ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectV1Tags struct {
	ID   string `json:"id,omitempty"`   // UUID of tag
	Name string `json:"name,omitempty"` // Name of tag
}
type ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectV1Templates struct {
	Tags                    *[]ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesTags                   `json:"tags,omitempty"`                    //
	Author                  string                                                                                        `json:"author,omitempty"`                  // Author of template
	Composite               *bool                                                                                         `json:"composite,omitempty"`               // Is it composite template
	ContainingTemplates     *[]ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesContainingTemplates    `json:"containingTemplates,omitempty"`     //
	CreateTime              *int                                                                                          `json:"createTime,omitempty"`              // Create time of template
	CustomParamsOrder       *bool                                                                                         `json:"customParamsOrder,omitempty"`       // Custom Params Order
	Description             string                                                                                        `json:"description,omitempty"`             // Description of template
	DeviceTypes             *[]ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesDeviceTypes            `json:"deviceTypes,omitempty"`             //
	FailurePolicy           string                                                                                        `json:"failurePolicy,omitempty"`           // Define failure policy if template provisioning fails
	ID                      string                                                                                        `json:"id,omitempty"`                      // UUID of template
	Language                string                                                                                        `json:"language,omitempty"`                // Template language (JINJA or VELOCITY)
	LastUpdateTime          *int                                                                                          `json:"lastUpdateTime,omitempty"`          // Update time of template
	LatestVersionTime       *int                                                                                          `json:"latestVersionTime,omitempty"`       // Latest versioned template time
	Name                    string                                                                                        `json:"name,omitempty"`                    // Name of template
	ParentTemplateID        string                                                                                        `json:"parentTemplateId,omitempty"`        // Parent templateID
	ProjectID               string                                                                                        `json:"projectId,omitempty"`               // Project UUID
	ProjectName             string                                                                                        `json:"projectName,omitempty"`             // Project name
	RollbackTemplateContent string                                                                                        `json:"rollbackTemplateContent,omitempty"` // Rollback template content
	RollbackTemplateParams  *[]ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesRollbackTemplateParams `json:"rollbackTemplateParams,omitempty"`  //
	SoftwareType            string                                                                                        `json:"softwareType,omitempty"`            // Applicable device software type
	SoftwareVariant         string                                                                                        `json:"softwareVariant,omitempty"`         // Applicable device software variant
	SoftwareVersion         string                                                                                        `json:"softwareVersion,omitempty"`         // Applicable device software version
	TemplateContent         string                                                                                        `json:"templateContent,omitempty"`         // Template content
	TemplateParams          *[]ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesTemplateParams         `json:"templateParams,omitempty"`          //
	ValidationErrors        *ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesValidationErrors         `json:"validationErrors,omitempty"`        //
	Version                 string                                                                                        `json:"version,omitempty"`                 // Current version of template
}
type ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesTags struct {
	ID   string `json:"id,omitempty"`   // UUID of tag
	Name string `json:"name,omitempty"` // Name of tag
}
type ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesContainingTemplates struct {
	Tags                   *[]ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesContainingTemplatesTags                   `json:"tags,omitempty"`                   //
	Composite              *bool                                                                                                            `json:"composite,omitempty"`              // Is it composite template
	Description            string                                                                                                           `json:"description,omitempty"`            // Description of template
	DeviceTypes            *[]ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesContainingTemplatesDeviceTypes            `json:"deviceTypes,omitempty"`            //
	ID                     string                                                                                                           `json:"id,omitempty"`                     // UUID of template
	Language               string                                                                                                           `json:"language,omitempty"`               // Template language (JINJA or VELOCITY)
	Name                   string                                                                                                           `json:"name,omitempty"`                   // Name of template
	ProjectName            string                                                                                                           `json:"projectName,omitempty"`            // Project name
	RollbackTemplateParams *[]ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesContainingTemplatesRollbackTemplateParams `json:"rollbackTemplateParams,omitempty"` //
	TemplateContent        string                                                                                                           `json:"templateContent,omitempty"`        // Template content
	TemplateParams         *[]ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesContainingTemplatesTemplateParams         `json:"templateParams,omitempty"`         //
	Version                string                                                                                                           `json:"version,omitempty"`                // Current version of template
}
type ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesContainingTemplatesTags struct {
	ID   string `json:"id,omitempty"`   // UUID of tag
	Name string `json:"name,omitempty"` // Name of tag
}
type ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesContainingTemplatesDeviceTypes struct {
	ProductFamily string `json:"productFamily,omitempty"` // Device family
	ProductSeries string `json:"productSeries,omitempty"` // Device series
	ProductType   string `json:"productType,omitempty"`   // Device type
}
type ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesContainingTemplatesRollbackTemplateParams struct {
	Binding         string                                                                                                                  `json:"binding,omitempty"`         // Bind to source
	CustomOrder     *int                                                                                                                    `json:"customOrder,omitempty"`     // CustomOrder of template param
	DataType        string                                                                                                                  `json:"dataType,omitempty"`        // Datatype of template param
	DefaultValue    string                                                                                                                  `json:"defaultValue,omitempty"`    // Default value of template param
	Description     string                                                                                                                  `json:"description,omitempty"`     // Description of template param
	DisplayName     string                                                                                                                  `json:"displayName,omitempty"`     // Display name of param
	Group           string                                                                                                                  `json:"group,omitempty"`           // group
	ID              string                                                                                                                  `json:"id,omitempty"`              // UUID of template param
	InstructionText string                                                                                                                  `json:"instructionText,omitempty"` // Instruction text for param
	Key             string                                                                                                                  `json:"key,omitempty"`             // key
	NotParam        *bool                                                                                                                   `json:"notParam,omitempty"`        // Is it not a variable
	Order           *int                                                                                                                    `json:"order,omitempty"`           // Order of template param
	ParamArray      *bool                                                                                                                   `json:"paramArray,omitempty"`      // Is it an array
	ParameterName   string                                                                                                                  `json:"parameterName,omitempty"`   // Name of template param
	Provider        string                                                                                                                  `json:"provider,omitempty"`        // provider
	Range           *[]ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesContainingTemplatesRollbackTemplateParamsRange   `json:"range,omitempty"`           //
	Required        *bool                                                                                                                   `json:"required,omitempty"`        // Is param required
	Selection       *ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesContainingTemplatesRollbackTemplateParamsSelection `json:"selection,omitempty"`       //
}
type ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesContainingTemplatesRollbackTemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       // UUID of range
	MaxValue *int   `json:"maxValue,omitempty"` // Max value of range
	MinValue *int   `json:"minValue,omitempty"` // Min value of range
}
type ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesContainingTemplatesRollbackTemplateParamsSelection struct {
	DefaultSelectedValues []string                                                                                                                               `json:"defaultSelectedValues,omitempty"` // Default selection values
	ID                    string                                                                                                                                 `json:"id,omitempty"`                    // UUID of selection
	SelectionType         string                                                                                                                                 `json:"selectionType,omitempty"`         // Type of selection(SINGLE_SELECT or MULTI_SELECT)
	SelectionValues       *ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesContainingTemplatesRollbackTemplateParamsSelectionSelectionValues `json:"selectionValues,omitempty"`       // Selection values
}
type ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesContainingTemplatesRollbackTemplateParamsSelectionSelectionValues interface{}
type ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesContainingTemplatesTemplateParams struct {
	Binding         string                                                                                                          `json:"binding,omitempty"`         // Bind to source
	CustomOrder     *int                                                                                                            `json:"customOrder,omitempty"`     // CustomOrder of template param
	DataType        string                                                                                                          `json:"dataType,omitempty"`        // Datatype of template param
	DefaultValue    string                                                                                                          `json:"defaultValue,omitempty"`    // Default value of template param
	Description     string                                                                                                          `json:"description,omitempty"`     // Description of template param
	DisplayName     string                                                                                                          `json:"displayName,omitempty"`     // Display name of param
	Group           string                                                                                                          `json:"group,omitempty"`           // group
	ID              string                                                                                                          `json:"id,omitempty"`              // UUID of template param
	InstructionText string                                                                                                          `json:"instructionText,omitempty"` // Instruction text for param
	Key             string                                                                                                          `json:"key,omitempty"`             // key
	NotParam        *bool                                                                                                           `json:"notParam,omitempty"`        // Is it not a variable
	Order           *int                                                                                                            `json:"order,omitempty"`           // Order of template param
	ParamArray      *bool                                                                                                           `json:"paramArray,omitempty"`      // Is it an array
	ParameterName   string                                                                                                          `json:"parameterName,omitempty"`   // Name of template param
	Provider        string                                                                                                          `json:"provider,omitempty"`        // provider
	Range           *[]ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesContainingTemplatesTemplateParamsRange   `json:"range,omitempty"`           //
	Required        *bool                                                                                                           `json:"required,omitempty"`        // Is param required
	Selection       *ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesContainingTemplatesTemplateParamsSelection `json:"selection,omitempty"`       //
}
type ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesContainingTemplatesTemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       // UUID of range
	MaxValue *int   `json:"maxValue,omitempty"` // Max value of range
	MinValue *int   `json:"minValue,omitempty"` // Min value of range
}
type ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesContainingTemplatesTemplateParamsSelection struct {
	DefaultSelectedValues []string                                                                                                                       `json:"defaultSelectedValues,omitempty"` // Default selection values
	ID                    string                                                                                                                         `json:"id,omitempty"`                    // UUID of selection
	SelectionType         string                                                                                                                         `json:"selectionType,omitempty"`         // Type of selection(SINGLE_SELECT or MULTI_SELECT)
	SelectionValues       *ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesContainingTemplatesTemplateParamsSelectionSelectionValues `json:"selectionValues,omitempty"`       // Selection values
}
type ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesContainingTemplatesTemplateParamsSelectionSelectionValues interface{}
type ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesDeviceTypes struct {
	ProductFamily string `json:"productFamily,omitempty"` // Device family
	ProductSeries string `json:"productSeries,omitempty"` // Device series
	ProductType   string `json:"productType,omitempty"`   // Device type
}
type ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesRollbackTemplateParams struct {
	Binding         string                                                                                               `json:"binding,omitempty"`         // Bind to source
	CustomOrder     *int                                                                                                 `json:"customOrder,omitempty"`     // CustomOrder of template param
	DataType        string                                                                                               `json:"dataType,omitempty"`        // Datatype of template param
	DefaultValue    string                                                                                               `json:"defaultValue,omitempty"`    // Default value of template param
	Description     string                                                                                               `json:"description,omitempty"`     // Description of template param
	DisplayName     string                                                                                               `json:"displayName,omitempty"`     // Display name of param
	Group           string                                                                                               `json:"group,omitempty"`           // group
	ID              string                                                                                               `json:"id,omitempty"`              // UUID of template param
	InstructionText string                                                                                               `json:"instructionText,omitempty"` // Instruction text for param
	Key             string                                                                                               `json:"key,omitempty"`             // key
	NotParam        *bool                                                                                                `json:"notParam,omitempty"`        // Is it not a variable
	Order           *int                                                                                                 `json:"order,omitempty"`           // Order of template param
	ParamArray      *bool                                                                                                `json:"paramArray,omitempty"`      // Is it an array
	ParameterName   string                                                                                               `json:"parameterName,omitempty"`   // Name of template param
	Provider        string                                                                                               `json:"provider,omitempty"`        // provider
	Range           *[]ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesRollbackTemplateParamsRange   `json:"range,omitempty"`           //
	Required        *bool                                                                                                `json:"required,omitempty"`        // Is param required
	Selection       *ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesRollbackTemplateParamsSelection `json:"selection,omitempty"`       //
}
type ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesRollbackTemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       // UUID of range
	MaxValue *int   `json:"maxValue,omitempty"` // Max value of range
	MinValue *int   `json:"minValue,omitempty"` // Min value of range
}
type ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesRollbackTemplateParamsSelection struct {
	DefaultSelectedValues []string                                                                                                            `json:"defaultSelectedValues,omitempty"` // Default selection values
	ID                    string                                                                                                              `json:"id,omitempty"`                    // UUID of selection
	SelectionType         string                                                                                                              `json:"selectionType,omitempty"`         // Type of selection(SINGLE_SELECT or MULTI_SELECT)
	SelectionValues       *ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesRollbackTemplateParamsSelectionSelectionValues `json:"selectionValues,omitempty"`       // Selection values
}
type ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesRollbackTemplateParamsSelectionSelectionValues interface{}
type ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesTemplateParams struct {
	Binding         string                                                                                       `json:"binding,omitempty"`         // Bind to source
	CustomOrder     *int                                                                                         `json:"customOrder,omitempty"`     // CustomOrder of template param
	DataType        string                                                                                       `json:"dataType,omitempty"`        // Datatype of template param
	DefaultValue    string                                                                                       `json:"defaultValue,omitempty"`    // Default value of template param
	Description     string                                                                                       `json:"description,omitempty"`     // Description of template param
	DisplayName     string                                                                                       `json:"displayName,omitempty"`     // Display name of param
	Group           string                                                                                       `json:"group,omitempty"`           // group
	ID              string                                                                                       `json:"id,omitempty"`              // UUID of template param
	InstructionText string                                                                                       `json:"instructionText,omitempty"` // Instruction text for param
	Key             string                                                                                       `json:"key,omitempty"`             // key
	NotParam        *bool                                                                                        `json:"notParam,omitempty"`        // Is it not a variable
	Order           *int                                                                                         `json:"order,omitempty"`           // Order of template param
	ParamArray      *bool                                                                                        `json:"paramArray,omitempty"`      // Is it an array
	ParameterName   string                                                                                       `json:"parameterName,omitempty"`   // Name of template param
	Provider        string                                                                                       `json:"provider,omitempty"`        // provider
	Range           *[]ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesTemplateParamsRange   `json:"range,omitempty"`           //
	Required        *bool                                                                                        `json:"required,omitempty"`        // Is param required
	Selection       *ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesTemplateParamsSelection `json:"selection,omitempty"`       //
}
type ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesTemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       // UUID of range
	MaxValue *int   `json:"maxValue,omitempty"` // Max value of range
	MinValue *int   `json:"minValue,omitempty"` // Min value of range
}
type ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesTemplateParamsSelection struct {
	DefaultSelectedValues []string                                                                                                    `json:"defaultSelectedValues,omitempty"` // Default selection values
	ID                    string                                                                                                      `json:"id,omitempty"`                    // UUID of selection
	SelectionType         string                                                                                                      `json:"selectionType,omitempty"`         // Type of selection(SINGLE_SELECT or MULTI_SELECT)
	SelectionValues       *ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesTemplateParamsSelectionSelectionValues `json:"selectionValues,omitempty"`       // Selection values
}
type ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesTemplateParamsSelectionSelectionValues interface{}
type ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesValidationErrors struct {
	RollbackTemplateErrors *ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesValidationErrorsRollbackTemplateErrors `json:"rollbackTemplateErrors,omitempty"` // Validation or design conflicts errors of rollback template
	TemplateErrors         *ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesValidationErrorsTemplateErrors         `json:"templateErrors,omitempty"`         // Validation or design conflicts errors
	TemplateID             string                                                                                                      `json:"templateId,omitempty"`             // UUID of template
	TemplateVersion        string                                                                                                      `json:"templateVersion,omitempty"`        // Current version of template
}
type ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesValidationErrorsRollbackTemplateErrors interface{}
type ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesValidationErrorsTemplateErrors interface{}
type ResponseConfigurationTemplatesDeletesTheProjectV1 struct {
	Response *ResponseConfigurationTemplatesDeletesTheProjectResponse `json:"response,omitempty"` //
	Version  string                                                   `json:"version,omitempty"`  //
}
type ResponseConfigurationTemplatesDeletesTheProjectResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseConfigurationTemplatesCreateTemplateV1 struct {
	Response *ResponseConfigurationTemplatesCreateTemplateV1Response `json:"response,omitempty"` //
	Version  string                                                  `json:"version,omitempty"`  //
}
type ResponseConfigurationTemplatesCreateTemplateV1Response struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseConfigurationTemplatesGetsTheTemplatesAvailableV1 []ResponseItemConfigurationTemplatesGetsTheTemplatesAvailableV1 // Array of ResponseConfigurationTemplatesGetsTheTemplatesAvailableV1
type ResponseItemConfigurationTemplatesGetsTheTemplatesAvailableV1 struct {
	Composite    *bool                                                                        `json:"composite,omitempty"`    // Is it composite template
	Name         string                                                                       `json:"name,omitempty"`         // Name of template
	ProjectID    string                                                                       `json:"projectId,omitempty"`    // UUID of project
	ProjectName  string                                                                       `json:"projectName,omitempty"`  // Name of project
	TemplateID   string                                                                       `json:"templateId,omitempty"`   // UUID of template
	VersionsInfo *[]ResponseItemConfigurationTemplatesGetsTheTemplatesAvailableV1VersionsInfo `json:"versionsInfo,omitempty"` //
}
type ResponseItemConfigurationTemplatesGetsTheTemplatesAvailableV1VersionsInfo struct {
	Author         string `json:"author,omitempty"`         // Author of version template
	Description    string `json:"description,omitempty"`    // Description of template
	ID             string `json:"id,omitempty"`             // UUID of template
	Version        string `json:"version,omitempty"`        // Current version of template
	VersionComment string `json:"versionComment,omitempty"` // Version comment
	VersionTime    *int   `json:"versionTime,omitempty"`    // Template version time
}
type ResponseConfigurationTemplatesUpdateTemplateV1 struct {
	Response *ResponseConfigurationTemplatesUpdateTemplateV1Response `json:"response,omitempty"` //
	Version  string                                                  `json:"version,omitempty"`  //
}
type ResponseConfigurationTemplatesUpdateTemplateV1Response struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseConfigurationTemplatesDeployTemplateV1 struct {
	DeploymentID    string                                                   `json:"deploymentId,omitempty"`    // UUID of deployment
	DeploymentName  string                                                   `json:"deploymentName,omitempty"`  // Name of deployment
	Devices         *[]ResponseConfigurationTemplatesDeployTemplateV1Devices `json:"devices,omitempty"`         //
	Duration        string                                                   `json:"duration,omitempty"`        // Total deployment duration
	EndTime         string                                                   `json:"endTime,omitempty"`         // Deployment end time
	ProjectName     string                                                   `json:"projectName,omitempty"`     // Name of project
	StartTime       string                                                   `json:"startTime,omitempty"`       // Deployment start time
	Status          string                                                   `json:"status,omitempty"`          // Current status of deployment
	StatusMessage   string                                                   `json:"statusMessage,omitempty"`   // Status message of deployment
	TemplateName    string                                                   `json:"templateName,omitempty"`    // Name of template deployed
	TemplateVersion string                                                   `json:"templateVersion,omitempty"` // Version of template deployed
}
type ResponseConfigurationTemplatesDeployTemplateV1Devices struct {
	DetailedStatusMessage string `json:"detailedStatusMessage,omitempty"` // Device detailed status message
	DeviceID              string `json:"deviceId,omitempty"`              // UUID of device
	Duration              string `json:"duration,omitempty"`              // Total duration of deployment
	EndTime               string `json:"endTime,omitempty"`               // EndTime of deployment
	IDentifier            string `json:"identifier,omitempty"`            // Identifier of device based on the target type
	IPAddress             string `json:"ipAddress,omitempty"`             // Device IPAddress
	Name                  string `json:"name,omitempty"`                  // Name of device
	StartTime             string `json:"startTime,omitempty"`             // StartTime of deployment
	Status                string `json:"status,omitempty"`                // Current status of device
	TargetType            string `json:"targetType,omitempty"`            // Target type of device
}
type ResponseConfigurationTemplatesStatusOfTemplateDeploymentV1 struct {
	DeploymentID    string                                                               `json:"deploymentId,omitempty"`    // UUID of deployment
	DeploymentName  string                                                               `json:"deploymentName,omitempty"`  // Name of deployment
	Devices         *[]ResponseConfigurationTemplatesStatusOfTemplateDeploymentV1Devices `json:"devices,omitempty"`         //
	Duration        string                                                               `json:"duration,omitempty"`        // Total deployment duration
	EndTime         string                                                               `json:"endTime,omitempty"`         // Deployment end time
	ProjectName     string                                                               `json:"projectName,omitempty"`     // Name of project
	StartTime       string                                                               `json:"startTime,omitempty"`       // Deployment start time
	Status          string                                                               `json:"status,omitempty"`          // Current status of deployment
	StatusMessage   string                                                               `json:"statusMessage,omitempty"`   // Status message of deployment
	TemplateName    string                                                               `json:"templateName,omitempty"`    // Name of template deployed
	TemplateVersion string                                                               `json:"templateVersion,omitempty"` // Version of template deployed
}
type ResponseConfigurationTemplatesStatusOfTemplateDeploymentV1Devices struct {
	DetailedStatusMessage string `json:"detailedStatusMessage,omitempty"` // Device detailed status message
	DeviceID              string `json:"deviceId,omitempty"`              // UUID of device
	Duration              string `json:"duration,omitempty"`              // Total duration of deployment
	EndTime               string `json:"endTime,omitempty"`               // EndTime of deployment
	IDentifier            string `json:"identifier,omitempty"`            // Identifier of device based on the target type
	IPAddress             string `json:"ipAddress,omitempty"`             // Device IPAddress
	Name                  string `json:"name,omitempty"`                  // Name of device
	StartTime             string `json:"startTime,omitempty"`             // StartTime of deployment
	Status                string `json:"status,omitempty"`                // Current status of device
	TargetType            string `json:"targetType,omitempty"`            // Target type of device
}
type ResponseConfigurationTemplatesExportsTheTemplatesForAGivenCriteriaV1 struct {
	Response *ResponseConfigurationTemplatesExportsTheTemplatesForAGivenCriteriaV1Response `json:"response,omitempty"` //
	Version  string                                                                        `json:"version,omitempty"`  //
}
type ResponseConfigurationTemplatesExportsTheTemplatesForAGivenCriteriaV1Response struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseConfigurationTemplatesPreviewTemplateV1 struct {
	CliPreview       string                                                           `json:"cliPreview,omitempty"`       // Generated template preview
	DeviceID         string                                                           `json:"deviceId,omitempty"`         // UUID of device
	TemplateID       string                                                           `json:"templateId,omitempty"`       // UUID of template
	ValidationErrors *ResponseConfigurationTemplatesPreviewTemplateV1ValidationErrors `json:"validationErrors,omitempty"` // Validation error in template content if any
}
type ResponseConfigurationTemplatesPreviewTemplateV1ValidationErrors interface{}
type ResponseConfigurationTemplatesVersionTemplateV1 struct {
	Response *ResponseConfigurationTemplatesVersionTemplateV1Response `json:"response,omitempty"` //
	Version  string                                                   `json:"version,omitempty"`  //
}
type ResponseConfigurationTemplatesVersionTemplateV1Response struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseConfigurationTemplatesGetsAllTheVersionsOfAGivenTemplateV1 []ResponseItemConfigurationTemplatesGetsAllTheVersionsOfAGivenTemplateV1 // Array of ResponseConfigurationTemplatesGetsAllTheVersionsOfAGivenTemplateV1
type ResponseItemConfigurationTemplatesGetsAllTheVersionsOfAGivenTemplateV1 struct {
	Composite    *bool                                                                                 `json:"composite,omitempty"`    // Is it composite template
	Name         string                                                                                `json:"name,omitempty"`         // Name of template
	ProjectID    string                                                                                `json:"projectId,omitempty"`    // UUID of project
	ProjectName  string                                                                                `json:"projectName,omitempty"`  // Name of project
	TemplateID   string                                                                                `json:"templateId,omitempty"`   // UUID of template
	VersionsInfo *[]ResponseItemConfigurationTemplatesGetsAllTheVersionsOfAGivenTemplateV1VersionsInfo `json:"versionsInfo,omitempty"` //
}
type ResponseItemConfigurationTemplatesGetsAllTheVersionsOfAGivenTemplateV1VersionsInfo struct {
	Author         string `json:"author,omitempty"`         // Author of version template
	Description    string `json:"description,omitempty"`    // Description of template
	ID             string `json:"id,omitempty"`             // UUID of template
	Version        string `json:"version,omitempty"`        // Current version of template
	VersionComment string `json:"versionComment,omitempty"` // Version comment
	VersionTime    *int   `json:"versionTime,omitempty"`    // Template version time
}
type ResponseConfigurationTemplatesDeletesTheTemplateV1 struct {
	Response *ResponseConfigurationTemplatesDeletesTheTemplateV1Response `json:"response,omitempty"` //
	Version  string                                                      `json:"version,omitempty"`  //
}
type ResponseConfigurationTemplatesDeletesTheTemplateV1Response struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateV1 struct {
	Tags                    *[]ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateV1Tags                   `json:"tags,omitempty"`                    //
	Author                  string                                                                               `json:"author,omitempty"`                  // Author of template
	Composite               *bool                                                                                `json:"composite,omitempty"`               // Is it composite template
	ContainingTemplates     *[]ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ContainingTemplates    `json:"containingTemplates,omitempty"`     //
	CreateTime              *int                                                                                 `json:"createTime,omitempty"`              // Create time of template
	CustomParamsOrder       *bool                                                                                `json:"customParamsOrder,omitempty"`       // Custom Params Order
	Description             string                                                                               `json:"description,omitempty"`             // Description of template
	DeviceTypes             *[]ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateV1DeviceTypes            `json:"deviceTypes,omitempty"`             //
	FailurePolicy           string                                                                               `json:"failurePolicy,omitempty"`           // Define failure policy if template provisioning fails
	ID                      string                                                                               `json:"id,omitempty"`                      // UUID of template
	Language                string                                                                               `json:"language,omitempty"`                // Template language (JINJA or VELOCITY)
	LastUpdateTime          *int                                                                                 `json:"lastUpdateTime,omitempty"`          // Update time of template
	LatestVersionTime       *int                                                                                 `json:"latestVersionTime,omitempty"`       // Latest versioned template time
	Name                    string                                                                               `json:"name,omitempty"`                    // Name of template
	ParentTemplateID        string                                                                               `json:"parentTemplateId,omitempty"`        // Parent templateID
	ProjectID               string                                                                               `json:"projectId,omitempty"`               // Project UUID
	ProjectName             string                                                                               `json:"projectName,omitempty"`             // Project name
	RollbackTemplateContent string                                                                               `json:"rollbackTemplateContent,omitempty"` // Rollback template content
	RollbackTemplateParams  *[]ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateV1RollbackTemplateParams `json:"rollbackTemplateParams,omitempty"`  //
	SoftwareType            string                                                                               `json:"softwareType,omitempty"`            // Applicable device software type
	SoftwareVariant         string                                                                               `json:"softwareVariant,omitempty"`         // Applicable device software variant
	SoftwareVersion         string                                                                               `json:"softwareVersion,omitempty"`         // Applicable device software version
	TemplateContent         string                                                                               `json:"templateContent,omitempty"`         // Template content
	TemplateParams          *[]ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateV1TemplateParams         `json:"templateParams,omitempty"`          //
	ValidationErrors        *ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ValidationErrors         `json:"validationErrors,omitempty"`        //
	Version                 string                                                                               `json:"version,omitempty"`                 // Current version of template
}
type ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateV1Tags struct {
	ID   string `json:"id,omitempty"`   // UUID of tag
	Name string `json:"name,omitempty"` // Name of tag
}
type ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ContainingTemplates struct {
	Tags                   *[]ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ContainingTemplatesTags                   `json:"tags,omitempty"`                   //
	Composite              *bool                                                                                                   `json:"composite,omitempty"`              // Is it composite template
	Description            string                                                                                                  `json:"description,omitempty"`            // Description of template
	DeviceTypes            *[]ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ContainingTemplatesDeviceTypes            `json:"deviceTypes,omitempty"`            //
	ID                     string                                                                                                  `json:"id,omitempty"`                     // UUID of template
	Language               string                                                                                                  `json:"language,omitempty"`               // Template language (JINJA or VELOCITY)
	Name                   string                                                                                                  `json:"name,omitempty"`                   // Name of template
	ProjectName            string                                                                                                  `json:"projectName,omitempty"`            // Project name
	RollbackTemplateParams *[]ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ContainingTemplatesRollbackTemplateParams `json:"rollbackTemplateParams,omitempty"` //
	TemplateContent        string                                                                                                  `json:"templateContent,omitempty"`        // Template content
	TemplateParams         *[]ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ContainingTemplatesTemplateParams         `json:"templateParams,omitempty"`         //
	Version                string                                                                                                  `json:"version,omitempty"`                // Current version of template
}
type ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ContainingTemplatesTags struct {
	ID   string `json:"id,omitempty"`   // UUID of tag
	Name string `json:"name,omitempty"` // Name of tag
}
type ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ContainingTemplatesDeviceTypes struct {
	ProductFamily string `json:"productFamily,omitempty"` // Device family
	ProductSeries string `json:"productSeries,omitempty"` // Device series
	ProductType   string `json:"productType,omitempty"`   // Device type
}
type ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ContainingTemplatesRollbackTemplateParams struct {
	Binding         string                                                                                                         `json:"binding,omitempty"`         // Bind to source
	CustomOrder     *int                                                                                                           `json:"customOrder,omitempty"`     // CustomOrder of template param
	DataType        string                                                                                                         `json:"dataType,omitempty"`        // Datatype of template param
	DefaultValue    string                                                                                                         `json:"defaultValue,omitempty"`    // Default value of template param
	Description     string                                                                                                         `json:"description,omitempty"`     // Description of template param
	DisplayName     string                                                                                                         `json:"displayName,omitempty"`     // Display name of param
	Group           string                                                                                                         `json:"group,omitempty"`           // group
	ID              string                                                                                                         `json:"id,omitempty"`              // UUID of template param
	InstructionText string                                                                                                         `json:"instructionText,omitempty"` // Instruction text for param
	Key             string                                                                                                         `json:"key,omitempty"`             // key
	NotParam        *bool                                                                                                          `json:"notParam,omitempty"`        // Is it not a variable
	Order           *int                                                                                                           `json:"order,omitempty"`           // Order of template param
	ParamArray      *bool                                                                                                          `json:"paramArray,omitempty"`      // Is it an array
	ParameterName   string                                                                                                         `json:"parameterName,omitempty"`   // Name of template param
	Provider        string                                                                                                         `json:"provider,omitempty"`        // provider
	Range           *[]ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ContainingTemplatesRollbackTemplateParamsRange   `json:"range,omitempty"`           //
	Required        *bool                                                                                                          `json:"required,omitempty"`        // Is param required
	Selection       *ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ContainingTemplatesRollbackTemplateParamsSelection `json:"selection,omitempty"`       //
}
type ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ContainingTemplatesRollbackTemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       // UUID of range
	MaxValue *int   `json:"maxValue,omitempty"` // Max value of range
	MinValue *int   `json:"minValue,omitempty"` // Min value of range
}
type ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ContainingTemplatesRollbackTemplateParamsSelection struct {
	DefaultSelectedValues []string                                                                                                                      `json:"defaultSelectedValues,omitempty"` // Default selection values
	ID                    string                                                                                                                        `json:"id,omitempty"`                    // UUID of selection
	SelectionType         string                                                                                                                        `json:"selectionType,omitempty"`         // Type of selection(SINGLE_SELECT or MULTI_SELECT)
	SelectionValues       *ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ContainingTemplatesRollbackTemplateParamsSelectionSelectionValues `json:"selectionValues,omitempty"`       // Selection values
}
type ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ContainingTemplatesRollbackTemplateParamsSelectionSelectionValues interface{}
type ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ContainingTemplatesTemplateParams struct {
	Binding         string                                                                                                 `json:"binding,omitempty"`         // Bind to source
	CustomOrder     *int                                                                                                   `json:"customOrder,omitempty"`     // CustomOrder of template param
	DataType        string                                                                                                 `json:"dataType,omitempty"`        // Datatype of template param
	DefaultValue    string                                                                                                 `json:"defaultValue,omitempty"`    // Default value of template param
	Description     string                                                                                                 `json:"description,omitempty"`     // Description of template param
	DisplayName     string                                                                                                 `json:"displayName,omitempty"`     // Display name of param
	Group           string                                                                                                 `json:"group,omitempty"`           // group
	ID              string                                                                                                 `json:"id,omitempty"`              // UUID of template param
	InstructionText string                                                                                                 `json:"instructionText,omitempty"` // Instruction text for param
	Key             string                                                                                                 `json:"key,omitempty"`             // key
	NotParam        *bool                                                                                                  `json:"notParam,omitempty"`        // Is it not a variable
	Order           *int                                                                                                   `json:"order,omitempty"`           // Order of template param
	ParamArray      *bool                                                                                                  `json:"paramArray,omitempty"`      // Is it an array
	ParameterName   string                                                                                                 `json:"parameterName,omitempty"`   // Name of template param
	Provider        string                                                                                                 `json:"provider,omitempty"`        // provider
	Range           *[]ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ContainingTemplatesTemplateParamsRange   `json:"range,omitempty"`           //
	Required        *bool                                                                                                  `json:"required,omitempty"`        // Is param required
	Selection       *ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ContainingTemplatesTemplateParamsSelection `json:"selection,omitempty"`       //
}
type ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ContainingTemplatesTemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       // UUID of range
	MaxValue *int   `json:"maxValue,omitempty"` // Max value of range
	MinValue *int   `json:"minValue,omitempty"` // Min value of range
}
type ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ContainingTemplatesTemplateParamsSelection struct {
	DefaultSelectedValues []string                                                                                                              `json:"defaultSelectedValues,omitempty"` // Default selection values
	ID                    string                                                                                                                `json:"id,omitempty"`                    // UUID of selection
	SelectionType         string                                                                                                                `json:"selectionType,omitempty"`         // Type of selection(SINGLE_SELECT or MULTI_SELECT)
	SelectionValues       *ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ContainingTemplatesTemplateParamsSelectionSelectionValues `json:"selectionValues,omitempty"`       // Selection values
}
type ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ContainingTemplatesTemplateParamsSelectionSelectionValues interface{}
type ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateV1DeviceTypes struct {
	ProductFamily string `json:"productFamily,omitempty"` // Device family
	ProductSeries string `json:"productSeries,omitempty"` // Device series
	ProductType   string `json:"productType,omitempty"`   // Device type
}
type ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateV1RollbackTemplateParams struct {
	Binding         string                                                                                      `json:"binding,omitempty"`         // Bind to source
	CustomOrder     *int                                                                                        `json:"customOrder,omitempty"`     // CustomOrder of template param
	DataType        string                                                                                      `json:"dataType,omitempty"`        // Datatype of template param
	DefaultValue    string                                                                                      `json:"defaultValue,omitempty"`    // Default value of template param
	Description     string                                                                                      `json:"description,omitempty"`     // Description of template param
	DisplayName     string                                                                                      `json:"displayName,omitempty"`     // Display name of param
	Group           string                                                                                      `json:"group,omitempty"`           // group
	ID              string                                                                                      `json:"id,omitempty"`              // UUID of template param
	InstructionText string                                                                                      `json:"instructionText,omitempty"` // Instruction text for param
	Key             string                                                                                      `json:"key,omitempty"`             // key
	NotParam        *bool                                                                                       `json:"notParam,omitempty"`        // Is it not a variable
	Order           *int                                                                                        `json:"order,omitempty"`           // Order of template param
	ParamArray      *bool                                                                                       `json:"paramArray,omitempty"`      // Is it an array
	ParameterName   string                                                                                      `json:"parameterName,omitempty"`   // Name of template param
	Provider        string                                                                                      `json:"provider,omitempty"`        // provider
	Range           *[]ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateV1RollbackTemplateParamsRange   `json:"range,omitempty"`           //
	Required        *bool                                                                                       `json:"required,omitempty"`        // Is param required
	Selection       *ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateV1RollbackTemplateParamsSelection `json:"selection,omitempty"`       //
}
type ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateV1RollbackTemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       // UUID of range
	MaxValue *int   `json:"maxValue,omitempty"` // Max value of range
	MinValue *int   `json:"minValue,omitempty"` // Min value of range
}
type ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateV1RollbackTemplateParamsSelection struct {
	DefaultSelectedValues []string                                                                                                   `json:"defaultSelectedValues,omitempty"` // Default selection values
	ID                    string                                                                                                     `json:"id,omitempty"`                    // UUID of selection
	SelectionType         string                                                                                                     `json:"selectionType,omitempty"`         // Type of selection(SINGLE_SELECT or MULTI_SELECT)
	SelectionValues       *ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateV1RollbackTemplateParamsSelectionSelectionValues `json:"selectionValues,omitempty"`       // Selection values
}
type ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateV1RollbackTemplateParamsSelectionSelectionValues interface{}
type ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateV1TemplateParams struct {
	Binding         string                                                                              `json:"binding,omitempty"`         // Bind to source
	CustomOrder     *int                                                                                `json:"customOrder,omitempty"`     // CustomOrder of template param
	DataType        string                                                                              `json:"dataType,omitempty"`        // Datatype of template param
	DefaultValue    string                                                                              `json:"defaultValue,omitempty"`    // Default value of template param
	Description     string                                                                              `json:"description,omitempty"`     // Description of template param
	DisplayName     string                                                                              `json:"displayName,omitempty"`     // Display name of param
	Group           string                                                                              `json:"group,omitempty"`           // group
	ID              string                                                                              `json:"id,omitempty"`              // UUID of template param
	InstructionText string                                                                              `json:"instructionText,omitempty"` // Instruction text for param
	Key             string                                                                              `json:"key,omitempty"`             // key
	NotParam        *bool                                                                               `json:"notParam,omitempty"`        // Is it not a variable
	Order           *int                                                                                `json:"order,omitempty"`           // Order of template param
	ParamArray      *bool                                                                               `json:"paramArray,omitempty"`      // Is it an array
	ParameterName   string                                                                              `json:"parameterName,omitempty"`   // Name of template param
	Provider        string                                                                              `json:"provider,omitempty"`        // provider
	Range           *[]ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateV1TemplateParamsRange   `json:"range,omitempty"`           //
	Required        *bool                                                                               `json:"required,omitempty"`        // Is param required
	Selection       *ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateV1TemplateParamsSelection `json:"selection,omitempty"`       //
}
type ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateV1TemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       // UUID of range
	MaxValue *int   `json:"maxValue,omitempty"` // Max value of range
	MinValue *int   `json:"minValue,omitempty"` // Min value of range
}
type ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateV1TemplateParamsSelection struct {
	DefaultSelectedValues []string                                                                                           `json:"defaultSelectedValues,omitempty"` // Default selection values
	ID                    string                                                                                             `json:"id,omitempty"`                    // UUID of selection
	SelectionType         string                                                                                             `json:"selectionType,omitempty"`         // Type of selection(SINGLE_SELECT or MULTI_SELECT)
	SelectionValues       *ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateV1TemplateParamsSelectionSelectionValues `json:"selectionValues,omitempty"`       // Selection values
}
type ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateV1TemplateParamsSelectionSelectionValues interface{}
type ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ValidationErrors struct {
	RollbackTemplateErrors *ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ValidationErrorsRollbackTemplateErrors `json:"rollbackTemplateErrors,omitempty"` // Validation or design conflicts errors of rollback template
	TemplateErrors         *ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ValidationErrorsTemplateErrors         `json:"templateErrors,omitempty"`         // Validation or design conflicts errors
	TemplateID             string                                                                                             `json:"templateId,omitempty"`             // UUID of template
	TemplateVersion        string                                                                                             `json:"templateVersion,omitempty"`        // Current version of template
}
type ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ValidationErrorsRollbackTemplateErrors interface{}
type ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ValidationErrorsTemplateErrors interface{}
type ResponseConfigurationTemplatesGetProjectsDetailsV2 struct {
	Response []ResponseConfigurationTemplatesGetProjectsDetailsResponse `json:"response,omitempty"` // Response
}

type ResponseConfigurationTemplatesGetProjectsDetailsResponse struct {
	CreateTime     *int                                                         `json:"createTime,omitempty"`     // Create time of project
	Description    string                                                       `json:"description,omitempty"`    // Description of project
	ID             string                                                       `json:"id,omitempty"`             // UUID of project
	IsDeletable    *bool                                                        `json:"isDeletable,omitempty"`    // Flag to check if project is deletable or not(for internal use only)
	LastUpdateTime *int                                                         `json:"lastUpdateTime,omitempty"` // Update time of project
	Name           string                                                       `json:"name,omitempty"`           // Name of project
	Tags           *[]ResponseConfigurationTemplatesGetProjectsDetailsV2Tags    `json:"tags,omitempty"`           //
	Templates      *ResponseConfigurationTemplatesGetProjectsDetailsV2Templates `json:"templates,omitempty"`      // List of templates within the project
}
type ResponseConfigurationTemplatesGetProjectsDetailsV2Tags struct {
	ID   string `json:"id,omitempty"`   // UUID of tag
	Name string `json:"name,omitempty"` // Name of tag
}
type ResponseConfigurationTemplatesGetProjectsDetailsV2Templates interface{}
type ResponseConfigurationTemplatesGetTemplatesV2Details struct {
	Response []ResponseConfigurationTemplatesGetTemplatesDetailsV2 `json:"response,omitempty"` // Response
}
type ResponseConfigurationTemplatesGetTemplatesDetailsV2 struct {
	Author                  string                                                                       `json:"author,omitempty"`                  // Author of template
	Composite               *bool                                                                        `json:"composite,omitempty"`               // Is it composite template
	ContainingTemplates     *[]ResponseConfigurationTemplatesGetTemplatesDetailsV2ContainingTemplates    `json:"containingTemplates,omitempty"`     //
	CreateTime              *int                                                                         `json:"createTime,omitempty"`              // Create time of template
	CustomParamsOrder       *bool                                                                        `json:"customParamsOrder,omitempty"`       // Custom Params Order
	Description             string                                                                       `json:"description,omitempty"`             // Description of template
	DeviceTypes             *[]ResponseConfigurationTemplatesGetTemplatesDetailsV2DeviceTypes            `json:"deviceTypes,omitempty"`             //
	FailurePolicy           string                                                                       `json:"failurePolicy,omitempty"`           // Define failure policy if template provisioning fails
	ID                      string                                                                       `json:"id,omitempty"`                      // UUID of template
	Language                string                                                                       `json:"language,omitempty"`                // Template language (JINJA or VELOCITY)
	LastUpdateTime          *int                                                                         `json:"lastUpdateTime,omitempty"`          // Update time of template
	LatestVersionTime       *int                                                                         `json:"latestVersionTime,omitempty"`       // Latest versioned template time
	Name                    string                                                                       `json:"name,omitempty"`                    // Name of template
	ParentTemplateID        string                                                                       `json:"parentTemplateId,omitempty"`        // Parent templateID
	ProjectAssociated       *bool                                                                        `json:"projectAssociated,omitempty"`       //
	ProjectID               string                                                                       `json:"projectId,omitempty"`               // Project UUID
	ProjectName             string                                                                       `json:"projectName,omitempty"`             // Project name
	RollbackTemplateContent string                                                                       `json:"rollbackTemplateContent,omitempty"` // Rollback template content
	RollbackTemplateParams  *[]ResponseConfigurationTemplatesGetTemplatesDetailsV2RollbackTemplateParams `json:"rollbackTemplateParams,omitempty"`  //
	SoftwareType            string                                                                       `json:"softwareType,omitempty"`            // Applicable device software type
	SoftwareVariant         string                                                                       `json:"softwareVariant,omitempty"`         // Applicable device software variant
	SoftwareVersion         string                                                                       `json:"softwareVersion,omitempty"`         // Applicable device software version
	Tags                    *[]ResponseConfigurationTemplatesGetTemplatesDetailsV2Tags                   `json:"tags,omitempty"`                    //
	TemplateContent         string                                                                       `json:"templateContent,omitempty"`         // Template content
	TemplateParams          *[]ResponseConfigurationTemplatesGetTemplatesDetailsV2TemplateParams         `json:"templateParams,omitempty"`          //
	ValidationErrors        *ResponseConfigurationTemplatesGetTemplatesDetailsV2ValidationErrors         `json:"validationErrors,omitempty"`        //
	Version                 string                                                                       `json:"version,omitempty"`                 // Current version of template
	VersionsInfo            *[]ResponseConfigurationTemplatesGetTemplatesDetailsV2VersionsInfo           `json:"versionsInfo,omitempty"`            //
}
type ResponseConfigurationTemplatesGetTemplatesDetailsV2ContainingTemplates struct {
	Composite              *bool                                                                                           `json:"composite,omitempty"`              // Is it composite template
	Description            string                                                                                          `json:"description,omitempty"`            // Description of template
	DeviceTypes            *[]ResponseConfigurationTemplatesGetTemplatesDetailsV2ContainingTemplatesDeviceTypes            `json:"deviceTypes,omitempty"`            //
	ID                     string                                                                                          `json:"id,omitempty"`                     // UUID of template
	Language               string                                                                                          `json:"language,omitempty"`               // Template language (JINJA or VELOCITY)
	Name                   string                                                                                          `json:"name,omitempty"`                   // Name of template
	ProjectName            string                                                                                          `json:"projectName,omitempty"`            // Project name
	RollbackTemplateParams *[]ResponseConfigurationTemplatesGetTemplatesDetailsV2ContainingTemplatesRollbackTemplateParams `json:"rollbackTemplateParams,omitempty"` //
	Tags                   *[]ResponseConfigurationTemplatesGetTemplatesDetailsV2ContainingTemplatesTags                   `json:"tags,omitempty"`                   //
	TemplateContent        string                                                                                          `json:"templateContent,omitempty"`        // Template content
	TemplateParams         *[]ResponseConfigurationTemplatesGetTemplatesDetailsV2ContainingTemplatesTemplateParams         `json:"templateParams,omitempty"`         //
	Version                string                                                                                          `json:"version,omitempty"`                // Current version of template
}
type ResponseConfigurationTemplatesGetTemplatesDetailsV2ContainingTemplatesDeviceTypes struct {
	ProductFamily string `json:"productFamily,omitempty"` // Device family
	ProductSeries string `json:"productSeries,omitempty"` // Device series
	ProductType   string `json:"productType,omitempty"`   // Device type
}
type ResponseConfigurationTemplatesGetTemplatesDetailsV2ContainingTemplatesRollbackTemplateParams struct {
	Binding         string                                                                                                 `json:"binding,omitempty"`         // Bind to source
	CustomOrder     *int                                                                                                   `json:"customOrder,omitempty"`     // CustomOrder of template param
	DataType        string                                                                                                 `json:"dataType,omitempty"`        // Datatype of template param
	DefaultValue    string                                                                                                 `json:"defaultValue,omitempty"`    // Default value of template param
	Description     string                                                                                                 `json:"description,omitempty"`     // Description of template param
	DisplayName     string                                                                                                 `json:"displayName,omitempty"`     // Display name of param
	Group           string                                                                                                 `json:"group,omitempty"`           // group
	ID              string                                                                                                 `json:"id,omitempty"`              // UUID of template param
	InstructionText string                                                                                                 `json:"instructionText,omitempty"` // Instruction text for param
	Key             string                                                                                                 `json:"key,omitempty"`             // key
	NotParam        *bool                                                                                                  `json:"notParam,omitempty"`        // Is it not a variable
	Order           *int                                                                                                   `json:"order,omitempty"`           // Order of template param
	ParamArray      *bool                                                                                                  `json:"paramArray,omitempty"`      // Is it an array
	ParameterName   string                                                                                                 `json:"parameterName,omitempty"`   // Name of template param
	Provider        string                                                                                                 `json:"provider,omitempty"`        // provider
	Range           *[]ResponseConfigurationTemplatesGetTemplatesDetailsV2ContainingTemplatesRollbackTemplateParamsRange   `json:"range,omitempty"`           //
	Required        *bool                                                                                                  `json:"required,omitempty"`        // Is param required
	Selection       *ResponseConfigurationTemplatesGetTemplatesDetailsV2ContainingTemplatesRollbackTemplateParamsSelection `json:"selection,omitempty"`       //
}
type ResponseConfigurationTemplatesGetTemplatesDetailsV2ContainingTemplatesRollbackTemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       // UUID of range
	MaxValue *int   `json:"maxValue,omitempty"` // Max value of range
	MinValue *int   `json:"minValue,omitempty"` // Min value of range
}
type ResponseConfigurationTemplatesGetTemplatesDetailsV2ContainingTemplatesRollbackTemplateParamsSelection struct {
	DefaultSelectedValues []string                                                                                                              `json:"defaultSelectedValues,omitempty"` // Default selection values
	ID                    string                                                                                                                `json:"id,omitempty"`                    // UUID of selection
	SelectionType         string                                                                                                                `json:"selectionType,omitempty"`         // Type of selection(SINGLE_SELECT or MULTI_SELECT)
	SelectionValues       *ResponseConfigurationTemplatesGetTemplatesDetailsV2ContainingTemplatesRollbackTemplateParamsSelectionSelectionValues `json:"selectionValues,omitempty"`       // Selection values
}
type ResponseConfigurationTemplatesGetTemplatesDetailsV2ContainingTemplatesRollbackTemplateParamsSelectionSelectionValues interface{}
type ResponseConfigurationTemplatesGetTemplatesDetailsV2ContainingTemplatesTags struct {
	ID   string `json:"id,omitempty"`   // UUID of tag
	Name string `json:"name,omitempty"` // Name of tag
}
type ResponseConfigurationTemplatesGetTemplatesDetailsV2ContainingTemplatesTemplateParams struct {
	Binding         string                                                                                         `json:"binding,omitempty"`         // Bind to source
	CustomOrder     *int                                                                                           `json:"customOrder,omitempty"`     // CustomOrder of template param
	DataType        string                                                                                         `json:"dataType,omitempty"`        // Datatype of template param
	DefaultValue    string                                                                                         `json:"defaultValue,omitempty"`    // Default value of template param
	Description     string                                                                                         `json:"description,omitempty"`     // Description of template param
	DisplayName     string                                                                                         `json:"displayName,omitempty"`     // Display name of param
	Group           string                                                                                         `json:"group,omitempty"`           // group
	ID              string                                                                                         `json:"id,omitempty"`              // UUID of template param
	InstructionText string                                                                                         `json:"instructionText,omitempty"` // Instruction text for param
	Key             string                                                                                         `json:"key,omitempty"`             // key
	NotParam        *bool                                                                                          `json:"notParam,omitempty"`        // Is it not a variable
	Order           *int                                                                                           `json:"order,omitempty"`           // Order of template param
	ParamArray      *bool                                                                                          `json:"paramArray,omitempty"`      // Is it an array
	ParameterName   string                                                                                         `json:"parameterName,omitempty"`   // Name of template param
	Provider        string                                                                                         `json:"provider,omitempty"`        // provider
	Range           *[]ResponseConfigurationTemplatesGetTemplatesDetailsV2ContainingTemplatesTemplateParamsRange   `json:"range,omitempty"`           //
	Required        *bool                                                                                          `json:"required,omitempty"`        // Is param required
	Selection       *ResponseConfigurationTemplatesGetTemplatesDetailsV2ContainingTemplatesTemplateParamsSelection `json:"selection,omitempty"`       //
}
type ResponseConfigurationTemplatesGetTemplatesDetailsV2ContainingTemplatesTemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       // UUID of range
	MaxValue *int   `json:"maxValue,omitempty"` // Max value of range
	MinValue *int   `json:"minValue,omitempty"` // Min value of range
}
type ResponseConfigurationTemplatesGetTemplatesDetailsV2ContainingTemplatesTemplateParamsSelection struct {
	DefaultSelectedValues []string                                                                                                      `json:"defaultSelectedValues,omitempty"` // Default selection values
	ID                    string                                                                                                        `json:"id,omitempty"`                    // UUID of selection
	SelectionType         string                                                                                                        `json:"selectionType,omitempty"`         // Type of selection(SINGLE_SELECT or MULTI_SELECT)
	SelectionValues       *ResponseConfigurationTemplatesGetTemplatesDetailsV2ContainingTemplatesTemplateParamsSelectionSelectionValues `json:"selectionValues,omitempty"`       // Selection values
}
type ResponseConfigurationTemplatesGetTemplatesDetailsV2ContainingTemplatesTemplateParamsSelectionSelectionValues interface{}
type ResponseConfigurationTemplatesGetTemplatesDetailsV2DeviceTypes struct {
	ProductFamily string `json:"productFamily,omitempty"` // Device family
	ProductSeries string `json:"productSeries,omitempty"` // Device series
	ProductType   string `json:"productType,omitempty"`   // Device type
}
type ResponseConfigurationTemplatesGetTemplatesDetailsV2RollbackTemplateParams struct {
	Binding         string                                                                              `json:"binding,omitempty"`         // Bind to source
	CustomOrder     *int                                                                                `json:"customOrder,omitempty"`     // CustomOrder of template param
	DataType        string                                                                              `json:"dataType,omitempty"`        // Datatype of template param
	DefaultValue    string                                                                              `json:"defaultValue,omitempty"`    // Default value of template param
	Description     string                                                                              `json:"description,omitempty"`     // Description of template param
	DisplayName     string                                                                              `json:"displayName,omitempty"`     // Display name of param
	Group           string                                                                              `json:"group,omitempty"`           // group
	ID              string                                                                              `json:"id,omitempty"`              // UUID of template param
	InstructionText string                                                                              `json:"instructionText,omitempty"` // Instruction text for param
	Key             string                                                                              `json:"key,omitempty"`             // key
	NotParam        *bool                                                                               `json:"notParam,omitempty"`        // Is it not a variable
	Order           *int                                                                                `json:"order,omitempty"`           // Order of template param
	ParamArray      *bool                                                                               `json:"paramArray,omitempty"`      // Is it an array
	ParameterName   string                                                                              `json:"parameterName,omitempty"`   // Name of template param
	Provider        string                                                                              `json:"provider,omitempty"`        // provider
	Range           *[]ResponseConfigurationTemplatesGetTemplatesDetailsV2RollbackTemplateParamsRange   `json:"range,omitempty"`           //
	Required        *bool                                                                               `json:"required,omitempty"`        // Is param required
	Selection       *ResponseConfigurationTemplatesGetTemplatesDetailsV2RollbackTemplateParamsSelection `json:"selection,omitempty"`       //
}
type ResponseConfigurationTemplatesGetTemplatesDetailsV2RollbackTemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       // UUID of range
	MaxValue *int   `json:"maxValue,omitempty"` // Max value of range
	MinValue *int   `json:"minValue,omitempty"` // Min value of range
}
type ResponseConfigurationTemplatesGetTemplatesDetailsV2RollbackTemplateParamsSelection struct {
	DefaultSelectedValues []string                                                                                           `json:"defaultSelectedValues,omitempty"` // Default selection values
	ID                    string                                                                                             `json:"id,omitempty"`                    // UUID of selection
	SelectionType         string                                                                                             `json:"selectionType,omitempty"`         // Type of selection(SINGLE_SELECT or MULTI_SELECT)
	SelectionValues       *ResponseConfigurationTemplatesGetTemplatesDetailsV2RollbackTemplateParamsSelectionSelectionValues `json:"selectionValues,omitempty"`       // Selection values
}
type ResponseConfigurationTemplatesGetTemplatesDetailsV2RollbackTemplateParamsSelectionSelectionValues interface{}
type ResponseConfigurationTemplatesGetTemplatesDetailsV2Tags struct {
	ID   string `json:"id,omitempty"`   // UUID of tag
	Name string `json:"name,omitempty"` // Name of tag
}
type ResponseConfigurationTemplatesGetTemplatesDetailsV2TemplateParams struct {
	Binding         string                                                                      `json:"binding,omitempty"`         // Bind to source
	CustomOrder     *int                                                                        `json:"customOrder,omitempty"`     // CustomOrder of template param
	DataType        string                                                                      `json:"dataType,omitempty"`        // Datatype of template param
	DefaultValue    string                                                                      `json:"defaultValue,omitempty"`    // Default value of template param
	Description     string                                                                      `json:"description,omitempty"`     // Description of template param
	DisplayName     string                                                                      `json:"displayName,omitempty"`     // Display name of param
	Group           string                                                                      `json:"group,omitempty"`           // group
	ID              string                                                                      `json:"id,omitempty"`              // UUID of template param
	InstructionText string                                                                      `json:"instructionText,omitempty"` // Instruction text for param
	Key             string                                                                      `json:"key,omitempty"`             // key
	NotParam        *bool                                                                       `json:"notParam,omitempty"`        // Is it not a variable
	Order           *int                                                                        `json:"order,omitempty"`           // Order of template param
	ParamArray      *bool                                                                       `json:"paramArray,omitempty"`      // Is it an array
	ParameterName   string                                                                      `json:"parameterName,omitempty"`   // Name of template param
	Provider        string                                                                      `json:"provider,omitempty"`        // provider
	Range           *[]ResponseConfigurationTemplatesGetTemplatesDetailsV2TemplateParamsRange   `json:"range,omitempty"`           //
	Required        *bool                                                                       `json:"required,omitempty"`        // Is param required
	Selection       *ResponseConfigurationTemplatesGetTemplatesDetailsV2TemplateParamsSelection `json:"selection,omitempty"`       //
}
type ResponseConfigurationTemplatesGetTemplatesDetailsV2TemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       // UUID of range
	MaxValue *int   `json:"maxValue,omitempty"` // Max value of range
	MinValue *int   `json:"minValue,omitempty"` // Min value of range
}
type ResponseConfigurationTemplatesGetTemplatesDetailsV2TemplateParamsSelection struct {
	DefaultSelectedValues []string                                                                                   `json:"defaultSelectedValues,omitempty"` // Default selection values
	ID                    string                                                                                     `json:"id,omitempty"`                    // UUID of selection
	SelectionType         string                                                                                     `json:"selectionType,omitempty"`         // Type of selection(SINGLE_SELECT or MULTI_SELECT)
	SelectionValues       *ResponseConfigurationTemplatesGetTemplatesDetailsV2TemplateParamsSelectionSelectionValues `json:"selectionValues,omitempty"`       // Selection values
}
type ResponseConfigurationTemplatesGetTemplatesDetailsV2TemplateParamsSelectionSelectionValues interface{}
type ResponseConfigurationTemplatesGetTemplatesDetailsV2ValidationErrors struct {
	RollbackTemplateErrors *ResponseConfigurationTemplatesGetTemplatesDetailsV2ValidationErrorsRollbackTemplateErrors `json:"rollbackTemplateErrors,omitempty"` // Validation or design conflicts errors of rollback template
	TemplateErrors         *ResponseConfigurationTemplatesGetTemplatesDetailsV2ValidationErrorsTemplateErrors         `json:"templateErrors,omitempty"`         // Validation or design conflicts errors
	TemplateID             string                                                                                     `json:"templateId,omitempty"`             // UUID of template
	TemplateVersion        string                                                                                     `json:"templateVersion,omitempty"`        // Current version of template
}
type ResponseConfigurationTemplatesGetTemplatesDetailsV2ValidationErrorsRollbackTemplateErrors interface{}
type ResponseConfigurationTemplatesGetTemplatesDetailsV2ValidationErrorsTemplateErrors interface{}
type ResponseConfigurationTemplatesGetTemplatesDetailsV2VersionsInfo struct {
	Author         string `json:"author,omitempty"`         //
	Description    string `json:"description,omitempty"`    //
	ID             string `json:"id,omitempty"`             //
	Version        string `json:"version,omitempty"`        //
	VersionComment string `json:"versionComment,omitempty"` //
	VersionTime    *int   `json:"versionTime,omitempty"`    //
}
type ResponseConfigurationTemplatesDeployTemplateV2 struct {
	Response *ResponseConfigurationTemplatesDeployTemplateV2Response `json:"response,omitempty"` //
	Version  string                                                  `json:"version,omitempty"`  //
}
type ResponseConfigurationTemplatesDeployTemplateV2Response struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type RequestConfigurationTemplatesCreateProjectV1 struct {
	Tags           *[]RequestConfigurationTemplatesCreateProjectV1Tags      `json:"tags,omitempty"`           //
	CreateTime     *int                                                     `json:"createTime,omitempty"`     // Create time of project
	Description    string                                                   `json:"description,omitempty"`    // Description of project
	ID             string                                                   `json:"id,omitempty"`             // UUID of project
	LastUpdateTime *int                                                     `json:"lastUpdateTime,omitempty"` // Update time of project
	Name           string                                                   `json:"name,omitempty"`           // Name of project
	Templates      *[]RequestConfigurationTemplatesCreateProjectV1Templates `json:"templates,omitempty"`      // List of templates within the project
}
type RequestConfigurationTemplatesCreateProjectV1Tags struct {
	ID   string `json:"id,omitempty"`   // UUID of tag
	Name string `json:"name,omitempty"` // Name of tag
}
type RequestConfigurationTemplatesCreateProjectV1Templates interface{}
type RequestConfigurationTemplatesUpdateProjectV1 struct {
	Tags           *[]RequestConfigurationTemplatesUpdateProjectV1Tags    `json:"tags,omitempty"`           //
	CreateTime     *int                                                   `json:"createTime,omitempty"`     // Create time of project
	Description    string                                                 `json:"description,omitempty"`    // Description of project
	ID             string                                                 `json:"id,omitempty"`             // UUID of project
	LastUpdateTime *int                                                   `json:"lastUpdateTime,omitempty"` // Update time of project
	Name           string                                                 `json:"name,omitempty"`           // Name of project
	Templates      *RequestConfigurationTemplatesUpdateProjectV1Templates `json:"templates,omitempty"`      // List of templates within the project
}
type RequestConfigurationTemplatesUpdateProjectV1Tags struct {
	ID   string `json:"id,omitempty"`   // UUID of tag
	Name string `json:"name,omitempty"` // Name of tag
}
type RequestConfigurationTemplatesUpdateProjectV1Templates interface{}                                                                           // # Review unknown case
type RequestConfigurationTemplatesExportsTheProjectsForAGivenCriteriaV1 []RequestItemConfigurationTemplatesExportsTheProjectsForAGivenCriteriaV1 // Array of RequestConfigurationTemplatesExportsTheProjectsForAGivenCriteriaV1
type RequestItemConfigurationTemplatesExportsTheProjectsForAGivenCriteriaV1 interface{}
type RequestConfigurationTemplatesImportsTheTemplatesProvidedV1 []RequestItemConfigurationTemplatesImportsTheTemplatesProvidedV1 // Array of RequestConfigurationTemplatesImportsTheTemplatesProvidedV1
type RequestItemConfigurationTemplatesImportsTheTemplatesProvidedV1 struct {
	Tags                    *[]RequestItemConfigurationTemplatesImportsTheTemplatesProvidedV1Tags                   `json:"tags,omitempty"`                    //
	Author                  string                                                                                  `json:"author,omitempty"`                  // Author of template
	Composite               *bool                                                                                   `json:"composite,omitempty"`               // Is it composite template
	ContainingTemplates     *[]RequestItemConfigurationTemplatesImportsTheTemplatesProvidedV1ContainingTemplates    `json:"containingTemplates,omitempty"`     //
	CreateTime              *int                                                                                    `json:"createTime,omitempty"`              // Create time of template
	CustomParamsOrder       *bool                                                                                   `json:"customParamsOrder,omitempty"`       // Custom Params Order
	Description             string                                                                                  `json:"description,omitempty"`             // Description of template
	DeviceTypes             *[]RequestItemConfigurationTemplatesImportsTheTemplatesProvidedV1DeviceTypes            `json:"deviceTypes,omitempty"`             //
	FailurePolicy           string                                                                                  `json:"failurePolicy,omitempty"`           // Define failure policy if template provisioning fails
	ID                      string                                                                                  `json:"id,omitempty"`                      // UUID of template
	Language                string                                                                                  `json:"language,omitempty"`                // Template language (JINJA or VELOCITY)
	LastUpdateTime          *int                                                                                    `json:"lastUpdateTime,omitempty"`          // Update time of template
	LatestVersionTime       *int                                                                                    `json:"latestVersionTime,omitempty"`       // Latest versioned template time
	Name                    string                                                                                  `json:"name,omitempty"`                    // Name of template
	ParentTemplateID        string                                                                                  `json:"parentTemplateId,omitempty"`        // Parent templateID
	ProjectID               string                                                                                  `json:"projectId,omitempty"`               // Project UUID
	ProjectName             string                                                                                  `json:"projectName,omitempty"`             // Project name
	RollbackTemplateContent string                                                                                  `json:"rollbackTemplateContent,omitempty"` // Rollback template content
	RollbackTemplateParams  *[]RequestItemConfigurationTemplatesImportsTheTemplatesProvidedV1RollbackTemplateParams `json:"rollbackTemplateParams,omitempty"`  //
	SoftwareType            string                                                                                  `json:"softwareType,omitempty"`            // Applicable device software type
	SoftwareVariant         string                                                                                  `json:"softwareVariant,omitempty"`         // Applicable device software variant
	SoftwareVersion         string                                                                                  `json:"softwareVersion,omitempty"`         // Applicable device software version
	TemplateContent         string                                                                                  `json:"templateContent,omitempty"`         // Template content
	TemplateParams          *[]RequestItemConfigurationTemplatesImportsTheTemplatesProvidedV1TemplateParams         `json:"templateParams,omitempty"`          //
	ValidationErrors        *RequestItemConfigurationTemplatesImportsTheTemplatesProvidedV1ValidationErrors         `json:"validationErrors,omitempty"`        //
	Version                 string                                                                                  `json:"version,omitempty"`                 // Current version of template
}
type RequestItemConfigurationTemplatesImportsTheTemplatesProvidedV1Tags struct {
	ID   string `json:"id,omitempty"`   // UUID of tag
	Name string `json:"name,omitempty"` // Name of tag
}
type RequestItemConfigurationTemplatesImportsTheTemplatesProvidedV1ContainingTemplates struct {
	Tags                   *[]RequestItemConfigurationTemplatesImportsTheTemplatesProvidedV1ContainingTemplatesTags                   `json:"tags,omitempty"`                   //
	Composite              *bool                                                                                                      `json:"composite,omitempty"`              // Is it composite template
	Description            string                                                                                                     `json:"description,omitempty"`            // Description of template
	DeviceTypes            *[]RequestItemConfigurationTemplatesImportsTheTemplatesProvidedV1ContainingTemplatesDeviceTypes            `json:"deviceTypes,omitempty"`            //
	ID                     string                                                                                                     `json:"id,omitempty"`                     // UUID of template
	Language               string                                                                                                     `json:"language,omitempty"`               // Template language (JINJA or VELOCITY)
	Name                   string                                                                                                     `json:"name,omitempty"`                   // Name of template
	ProjectName            string                                                                                                     `json:"projectName,omitempty"`            // Project name
	RollbackTemplateParams *[]RequestItemConfigurationTemplatesImportsTheTemplatesProvidedV1ContainingTemplatesRollbackTemplateParams `json:"rollbackTemplateParams,omitempty"` //
	TemplateContent        string                                                                                                     `json:"templateContent,omitempty"`        // Template content
	TemplateParams         *[]RequestItemConfigurationTemplatesImportsTheTemplatesProvidedV1ContainingTemplatesTemplateParams         `json:"templateParams,omitempty"`         //
	Version                string                                                                                                     `json:"version,omitempty"`                // Current version of template
}
type RequestItemConfigurationTemplatesImportsTheTemplatesProvidedV1ContainingTemplatesTags struct {
	ID   string `json:"id,omitempty"`   // UUID of tag
	Name string `json:"name,omitempty"` // Name of tag
}
type RequestItemConfigurationTemplatesImportsTheTemplatesProvidedV1ContainingTemplatesDeviceTypes struct {
	ProductFamily string `json:"productFamily,omitempty"` // Device family
	ProductSeries string `json:"productSeries,omitempty"` // Device series
	ProductType   string `json:"productType,omitempty"`   // Device type
}
type RequestItemConfigurationTemplatesImportsTheTemplatesProvidedV1ContainingTemplatesRollbackTemplateParams struct {
	Binding         string                                                                                                            `json:"binding,omitempty"`         // Bind to source
	CustomOrder     *int                                                                                                              `json:"customOrder,omitempty"`     // CustomOrder of template param
	DataType        string                                                                                                            `json:"dataType,omitempty"`        // Datatype of template param
	DefaultValue    string                                                                                                            `json:"defaultValue,omitempty"`    // Default value of template param
	Description     string                                                                                                            `json:"description,omitempty"`     // Description of template param
	DisplayName     string                                                                                                            `json:"displayName,omitempty"`     // Display name of param
	Group           string                                                                                                            `json:"group,omitempty"`           // group
	ID              string                                                                                                            `json:"id,omitempty"`              // UUID of template param
	InstructionText string                                                                                                            `json:"instructionText,omitempty"` // Instruction text for param
	Key             string                                                                                                            `json:"key,omitempty"`             // key
	NotParam        *bool                                                                                                             `json:"notParam,omitempty"`        // Is it not a variable
	Order           *int                                                                                                              `json:"order,omitempty"`           // Order of template param
	ParamArray      *bool                                                                                                             `json:"paramArray,omitempty"`      // Is it an array
	ParameterName   string                                                                                                            `json:"parameterName,omitempty"`   // Name of template param
	Provider        string                                                                                                            `json:"provider,omitempty"`        // provider
	Range           *[]RequestItemConfigurationTemplatesImportsTheTemplatesProvidedV1ContainingTemplatesRollbackTemplateParamsRange   `json:"range,omitempty"`           //
	Required        *bool                                                                                                             `json:"required,omitempty"`        // Is param required
	Selection       *RequestItemConfigurationTemplatesImportsTheTemplatesProvidedV1ContainingTemplatesRollbackTemplateParamsSelection `json:"selection,omitempty"`       //
}
type RequestItemConfigurationTemplatesImportsTheTemplatesProvidedV1ContainingTemplatesRollbackTemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       // UUID of range
	MaxValue *int   `json:"maxValue,omitempty"` // Max value of range
	MinValue *int   `json:"minValue,omitempty"` // Min value of range
}
type RequestItemConfigurationTemplatesImportsTheTemplatesProvidedV1ContainingTemplatesRollbackTemplateParamsSelection struct {
	DefaultSelectedValues []string                                                                                                                         `json:"defaultSelectedValues,omitempty"` // Default selection values
	ID                    string                                                                                                                           `json:"id,omitempty"`                    // UUID of selection
	SelectionType         string                                                                                                                           `json:"selectionType,omitempty"`         // Type of selection(SINGLE_SELECT or MULTI_SELECT)
	SelectionValues       *RequestItemConfigurationTemplatesImportsTheTemplatesProvidedV1ContainingTemplatesRollbackTemplateParamsSelectionSelectionValues `json:"selectionValues,omitempty"`       // Selection values
}
type RequestItemConfigurationTemplatesImportsTheTemplatesProvidedV1ContainingTemplatesRollbackTemplateParamsSelectionSelectionValues interface{}
type RequestItemConfigurationTemplatesImportsTheTemplatesProvidedV1ContainingTemplatesTemplateParams struct {
	Binding         string                                                                                                    `json:"binding,omitempty"`         // Bind to source
	CustomOrder     *int                                                                                                      `json:"customOrder,omitempty"`     // CustomOrder of template param
	DataType        string                                                                                                    `json:"dataType,omitempty"`        // Datatype of template param
	DefaultValue    string                                                                                                    `json:"defaultValue,omitempty"`    // Default value of template param
	Description     string                                                                                                    `json:"description,omitempty"`     // Description of template param
	DisplayName     string                                                                                                    `json:"displayName,omitempty"`     // Display name of param
	Group           string                                                                                                    `json:"group,omitempty"`           // group
	ID              string                                                                                                    `json:"id,omitempty"`              // UUID of template param
	InstructionText string                                                                                                    `json:"instructionText,omitempty"` // Instruction text for param
	Key             string                                                                                                    `json:"key,omitempty"`             // key
	NotParam        *bool                                                                                                     `json:"notParam,omitempty"`        // Is it not a variable
	Order           *int                                                                                                      `json:"order,omitempty"`           // Order of template param
	ParamArray      *bool                                                                                                     `json:"paramArray,omitempty"`      // Is it an array
	ParameterName   string                                                                                                    `json:"parameterName,omitempty"`   // Name of template param
	Provider        string                                                                                                    `json:"provider,omitempty"`        // provider
	Range           *[]RequestItemConfigurationTemplatesImportsTheTemplatesProvidedV1ContainingTemplatesTemplateParamsRange   `json:"range,omitempty"`           //
	Required        *bool                                                                                                     `json:"required,omitempty"`        // Is param required
	Selection       *RequestItemConfigurationTemplatesImportsTheTemplatesProvidedV1ContainingTemplatesTemplateParamsSelection `json:"selection,omitempty"`       //
}
type RequestItemConfigurationTemplatesImportsTheTemplatesProvidedV1ContainingTemplatesTemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       // UUID of range
	MaxValue *int   `json:"maxValue,omitempty"` // Max value of range
	MinValue *int   `json:"minValue,omitempty"` // Min value of range
}
type RequestItemConfigurationTemplatesImportsTheTemplatesProvidedV1ContainingTemplatesTemplateParamsSelection struct {
	DefaultSelectedValues []string                                                                                                                 `json:"defaultSelectedValues,omitempty"` // Default selection values
	ID                    string                                                                                                                   `json:"id,omitempty"`                    // UUID of selection
	SelectionType         string                                                                                                                   `json:"selectionType,omitempty"`         // Type of selection(SINGLE_SELECT or MULTI_SELECT)
	SelectionValues       *RequestItemConfigurationTemplatesImportsTheTemplatesProvidedV1ContainingTemplatesTemplateParamsSelectionSelectionValues `json:"selectionValues,omitempty"`       // Selection values
}
type RequestItemConfigurationTemplatesImportsTheTemplatesProvidedV1ContainingTemplatesTemplateParamsSelectionSelectionValues interface{}
type RequestItemConfigurationTemplatesImportsTheTemplatesProvidedV1DeviceTypes struct {
	ProductFamily string `json:"productFamily,omitempty"` // Device family
	ProductSeries string `json:"productSeries,omitempty"` // Device series
	ProductType   string `json:"productType,omitempty"`   // Device type
}
type RequestItemConfigurationTemplatesImportsTheTemplatesProvidedV1RollbackTemplateParams struct {
	Binding         string                                                                                       `json:"binding,omitempty"`         // Bind to source
	CustomOrder     *int                                                                                         `json:"customOrder,omitempty"`     // CustomOrder of template param
	DataType        string                                                                                       `json:"dataType,omitempty"`        // Datatype of template param
	DefaultValue    string                                                                                       `json:"defaultValue,omitempty"`    // Default value of template param
	Description     string                                                                                       `json:"description,omitempty"`     // Description of template param
	DisplayName     string                                                                                       `json:"displayName,omitempty"`     // Display name of param
	Group           string                                                                                       `json:"group,omitempty"`           // group
	ID              string                                                                                       `json:"id,omitempty"`              // UUID of template param
	InstructionText string                                                                                       `json:"instructionText,omitempty"` // Instruction text for param
	Key             string                                                                                       `json:"key,omitempty"`             // key
	NotParam        *bool                                                                                        `json:"notParam,omitempty"`        // Is it not a variable
	Order           *int                                                                                         `json:"order,omitempty"`           // Order of template param
	ParamArray      *bool                                                                                        `json:"paramArray,omitempty"`      // Is it an array
	ParameterName   string                                                                                       `json:"parameterName,omitempty"`   // Name of template param
	Provider        string                                                                                       `json:"provider,omitempty"`        // provider
	Range           *[]RequestItemConfigurationTemplatesImportsTheTemplatesProvidedRollbackTemplateParamsRange   `json:"range,omitempty"`           //
	Required        *bool                                                                                        `json:"required,omitempty"`        // Is param required
	Selection       *RequestItemConfigurationTemplatesImportsTheTemplatesProvidedRollbackTemplateParamsSelection `json:"selection,omitempty"`       //
}
type RequestItemConfigurationTemplatesImportsTheTemplatesProvidedRollbackTemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       // UUID of range
	MaxValue *int   `json:"maxValue,omitempty"` // Max value of range
	MinValue *int   `json:"minValue,omitempty"` // Min value of range
}
type RequestItemConfigurationTemplatesImportsTheTemplatesProvidedRollbackTemplateParamsSelection struct {
	DefaultSelectedValues []string                                                                                                    `json:"defaultSelectedValues,omitempty"` // Default selection values
	ID                    string                                                                                                      `json:"id,omitempty"`                    // UUID of selection
	SelectionType         string                                                                                                      `json:"selectionType,omitempty"`         // Type of selection(SINGLE_SELECT or MULTI_SELECT)
	SelectionValues       *RequestItemConfigurationTemplatesImportsTheTemplatesProvidedRollbackTemplateParamsSelectionSelectionValues `json:"selectionValues,omitempty"`       // Selection values
}
type RequestItemConfigurationTemplatesImportsTheTemplatesProvidedRollbackTemplateParamsSelectionSelectionValues interface{}
type RequestItemConfigurationTemplatesImportsTheTemplatesProvidedTemplateParams struct {
	Binding         string                                                                               `json:"binding,omitempty"`         // Bind to source
	CustomOrder     *int                                                                                 `json:"customOrder,omitempty"`     // CustomOrder of template param
	DataType        string                                                                               `json:"dataType,omitempty"`        // Datatype of template param
	DefaultValue    string                                                                               `json:"defaultValue,omitempty"`    // Default value of template param
	Description     string                                                                               `json:"description,omitempty"`     // Description of template param
	DisplayName     string                                                                               `json:"displayName,omitempty"`     // Display name of param
	Group           string                                                                               `json:"group,omitempty"`           // group
	ID              string                                                                               `json:"id,omitempty"`              // UUID of template param
	InstructionText string                                                                               `json:"instructionText,omitempty"` // Instruction text for param
	Key             string                                                                               `json:"key,omitempty"`             // key
	NotParam        *bool                                                                                `json:"notParam,omitempty"`        // Is it not a variable
	Order           *int                                                                                 `json:"order,omitempty"`           // Order of template param
	ParamArray      *bool                                                                                `json:"paramArray,omitempty"`      // Is it an array
	ParameterName   string                                                                               `json:"parameterName,omitempty"`   // Name of template param
	Provider        string                                                                               `json:"provider,omitempty"`        // provider
	Range           *[]RequestItemConfigurationTemplatesImportsTheTemplatesProvidedTemplateParamsRange   `json:"range,omitempty"`           //
	Required        *bool                                                                                `json:"required,omitempty"`        // Is param required
	Selection       *RequestItemConfigurationTemplatesImportsTheTemplatesProvidedTemplateParamsSelection `json:"selection,omitempty"`       //
}
type RequestItemConfigurationTemplatesImportsTheTemplatesProvidedTemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       // UUID of range
	MaxValue *int   `json:"maxValue,omitempty"` // Max value of range
	MinValue *int   `json:"minValue,omitempty"` // Min value of range
}
type RequestItemConfigurationTemplatesImportsTheTemplatesProvidedTemplateParamsSelection struct {
	DefaultSelectedValues []string                                                                                            `json:"defaultSelectedValues,omitempty"` // Default selection values
	ID                    string                                                                                              `json:"id,omitempty"`                    // UUID of selection
	SelectionType         string                                                                                              `json:"selectionType,omitempty"`         // Type of selection(SINGLE_SELECT or MULTI_SELECT)
	SelectionValues       *RequestItemConfigurationTemplatesImportsTheTemplatesProvidedTemplateParamsSelectionSelectionValues `json:"selectionValues,omitempty"`       // Selection values
}
type RequestItemConfigurationTemplatesImportsTheTemplatesProvidedTemplateParamsSelectionSelectionValues interface{}
type RequestItemConfigurationTemplatesImportsTheTemplatesProvidedValidationErrors struct {
	RollbackTemplateErrors *RequestItemConfigurationTemplatesImportsTheTemplatesProvidedValidationErrorsRollbackTemplateErrors `json:"rollbackTemplateErrors,omitempty"` // Validation or design conflicts errors of rollback template
	TemplateErrors         *RequestItemConfigurationTemplatesImportsTheTemplatesProvidedValidationErrorsTemplateErrors         `json:"templateErrors,omitempty"`         // Validation or design conflicts errors
	TemplateID             string                                                                                              `json:"templateId,omitempty"`             // UUID of template
	TemplateVersion        string                                                                                              `json:"templateVersion,omitempty"`        // Current version of template
}
type RequestItemConfigurationTemplatesImportsTheTemplatesProvidedValidationErrorsRollbackTemplateErrors interface{}
type RequestItemConfigurationTemplatesImportsTheTemplatesProvidedValidationErrorsTemplateErrors interface{}
type RequestConfigurationTemplatesCreateTemplate struct {
	Tags                    *[]RequestConfigurationTemplatesCreateTemplateTags                     `json:"tags,omitempty"`                    //
	Author                  string                                                                 `json:"author,omitempty"`                  // Author of template
	Composite               *bool                                                                  `json:"composite,omitempty"`               // Is it composite template
	ContainingTemplates     *[]RequestConfigurationTemplatesCreateTemplateContainingTemplates      `json:"containingTemplates,omitempty"`     //
	CreateTime              *int                                                                   `json:"createTime,omitempty"`              // Create time of template
	CustomParamsOrder       *bool                                                                  `json:"customParamsOrder,omitempty"`       // Custom Params Order
	Description             string                                                                 `json:"description,omitempty"`             // Description of template
	DeviceTypes             *[]RequestConfigurationTemplatesCreateTemplateV1DeviceTypes            `json:"deviceTypes,omitempty"`             //
	FailurePolicy           string                                                                 `json:"failurePolicy,omitempty"`           // Define failure policy if template provisioning fails
	ID                      string                                                                 `json:"id,omitempty"`                      // UUID of template
	Language                string                                                                 `json:"language,omitempty"`                // Template language (JINJA or VELOCITY)
	LastUpdateTime          *int                                                                   `json:"lastUpdateTime,omitempty"`          // Update time of template
	LatestVersionTime       *int                                                                   `json:"latestVersionTime,omitempty"`       // Latest versioned template time
	Name                    string                                                                 `json:"name,omitempty"`                    // Name of template
	ParentTemplateID        string                                                                 `json:"parentTemplateId,omitempty"`        // Parent templateID
	ProjectID               string                                                                 `json:"projectId,omitempty"`               // Project UUID
	ProjectName             string                                                                 `json:"projectName,omitempty"`             // Project name
	RollbackTemplateContent string                                                                 `json:"rollbackTemplateContent,omitempty"` // Rollback template content
	RollbackTemplateParams  *[]RequestConfigurationTemplatesCreateTemplateV1RollbackTemplateParams `json:"rollbackTemplateParams,omitempty"`  //
	SoftwareType            string                                                                 `json:"softwareType,omitempty"`            // Applicable device software type
	SoftwareVariant         string                                                                 `json:"softwareVariant,omitempty"`         // Applicable device software variant
	SoftwareVersion         string                                                                 `json:"softwareVersion,omitempty"`         // Applicable device software version
	TemplateContent         string                                                                 `json:"templateContent,omitempty"`         // Template content
	TemplateParams          *[]RequestConfigurationTemplatesCreateTemplateV1TemplateParams         `json:"templateParams,omitempty"`          //
	ValidationErrors        *RequestConfigurationTemplatesCreateTemplateV1ValidationErrors         `json:"validationErrors,omitempty"`        //
	Version                 string                                                                 `json:"version,omitempty"`                 // Current version of template
}
type RequestConfigurationTemplatesCreateTemplateTags struct {
	ID   string `json:"id,omitempty"`   // UUID of tag
	Name string `json:"name,omitempty"` // Name of tag
}
type RequestConfigurationTemplatesCreateTemplateContainingTemplates struct {
	Tags                   *[]RequestConfigurationTemplatesCreateTemplateContainingTemplatesTags                   `json:"tags,omitempty"`                   //
	Composite              *bool                                                                                   `json:"composite,omitempty"`              // Is it composite template
	Description            string                                                                                  `json:"description,omitempty"`            // Description of template
	DeviceTypes            *[]RequestConfigurationTemplatesCreateTemplateContainingTemplatesDeviceTypes            `json:"deviceTypes,omitempty"`            //
	ID                     string                                                                                  `json:"id,omitempty"`                     // UUID of template
	Language               string                                                                                  `json:"language,omitempty"`               // Template language (JINJA or VELOCITY)
	Name                   string                                                                                  `json:"name,omitempty"`                   // Name of template
	ProjectName            string                                                                                  `json:"projectName,omitempty"`            // Project name
	RollbackTemplateParams *[]RequestConfigurationTemplatesCreateTemplateContainingTemplatesRollbackTemplateParams `json:"rollbackTemplateParams,omitempty"` //
	TemplateContent        string                                                                                  `json:"templateContent,omitempty"`        // Template content
	TemplateParams         *[]RequestConfigurationTemplatesCreateTemplateV1ContainingTemplatesTemplateParams       `json:"templateParams,omitempty"`         //
	Version                string                                                                                  `json:"version,omitempty"`                // Current version of template
}
type RequestConfigurationTemplatesCreateTemplateContainingTemplatesTags struct {
	ID   string `json:"id,omitempty"`   // UUID of tag
	Name string `json:"name,omitempty"` // Name of tag
}
type RequestConfigurationTemplatesCreateTemplateContainingTemplatesDeviceTypes struct {
	ProductFamily string `json:"productFamily,omitempty"` // Device family
	ProductSeries string `json:"productSeries,omitempty"` // Device series
	ProductType   string `json:"productType,omitempty"`   // Device type
}
type RequestConfigurationTemplatesCreateTemplateContainingTemplatesRollbackTemplateParams struct {
	Binding         string                                                                                         `json:"binding,omitempty"`         // Bind to source
	CustomOrder     *int                                                                                           `json:"customOrder,omitempty"`     // CustomOrder of template param
	DataType        string                                                                                         `json:"dataType,omitempty"`        // Datatype of template param
	DefaultValue    string                                                                                         `json:"defaultValue,omitempty"`    // Default value of template param
	Description     string                                                                                         `json:"description,omitempty"`     // Description of template param
	DisplayName     string                                                                                         `json:"displayName,omitempty"`     // Display name of param
	Group           string                                                                                         `json:"group,omitempty"`           // group
	ID              string                                                                                         `json:"id,omitempty"`              // UUID of template param
	InstructionText string                                                                                         `json:"instructionText,omitempty"` // Instruction text for param
	Key             string                                                                                         `json:"key,omitempty"`             // key
	NotParam        *bool                                                                                          `json:"notParam,omitempty"`        // Is it not a variable
	Order           *int                                                                                           `json:"order,omitempty"`           // Order of template param
	ParamArray      *bool                                                                                          `json:"paramArray,omitempty"`      // Is it an array
	ParameterName   string                                                                                         `json:"parameterName,omitempty"`   // Name of template param
	Provider        string                                                                                         `json:"provider,omitempty"`        // provider
	Range           *[]RequestItemConfigurationTemplatesImportsTheTemplatesProvidedV1RollbackTemplateParamsRange   `json:"range,omitempty"`           //
	Required        *bool                                                                                          `json:"required,omitempty"`        // Is param required
	Selection       *RequestItemConfigurationTemplatesImportsTheTemplatesProvidedV1RollbackTemplateParamsSelection `json:"selection,omitempty"`       //
}
type RequestItemConfigurationTemplatesImportsTheTemplatesProvidedV1RollbackTemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       // UUID of range
	MaxValue *int   `json:"maxValue,omitempty"` // Max value of range
	MinValue *int   `json:"minValue,omitempty"` // Min value of range
}
type RequestItemConfigurationTemplatesImportsTheTemplatesProvidedV1RollbackTemplateParamsSelection struct {
	DefaultSelectedValues []string                                                                                                      `json:"defaultSelectedValues,omitempty"` // Default selection values
	ID                    string                                                                                                        `json:"id,omitempty"`                    // UUID of selection
	SelectionType         string                                                                                                        `json:"selectionType,omitempty"`         // Type of selection(SINGLE_SELECT or MULTI_SELECT)
	SelectionValues       *RequestItemConfigurationTemplatesImportsTheTemplatesProvidedV1RollbackTemplateParamsSelectionSelectionValues `json:"selectionValues,omitempty"`       // Selection values
}
type RequestItemConfigurationTemplatesImportsTheTemplatesProvidedV1RollbackTemplateParamsSelectionSelectionValues interface{}
type RequestItemConfigurationTemplatesImportsTheTemplatesProvidedV1TemplateParams struct {
	Binding         string                                                                                 `json:"binding,omitempty"`         // Bind to source
	CustomOrder     *int                                                                                   `json:"customOrder,omitempty"`     // CustomOrder of template param
	DataType        string                                                                                 `json:"dataType,omitempty"`        // Datatype of template param
	DefaultValue    string                                                                                 `json:"defaultValue,omitempty"`    // Default value of template param
	Description     string                                                                                 `json:"description,omitempty"`     // Description of template param
	DisplayName     string                                                                                 `json:"displayName,omitempty"`     // Display name of param
	Group           string                                                                                 `json:"group,omitempty"`           // group
	ID              string                                                                                 `json:"id,omitempty"`              // UUID of template param
	InstructionText string                                                                                 `json:"instructionText,omitempty"` // Instruction text for param
	Key             string                                                                                 `json:"key,omitempty"`             // key
	NotParam        *bool                                                                                  `json:"notParam,omitempty"`        // Is it not a variable
	Order           *int                                                                                   `json:"order,omitempty"`           // Order of template param
	ParamArray      *bool                                                                                  `json:"paramArray,omitempty"`      // Is it an array
	ParameterName   string                                                                                 `json:"parameterName,omitempty"`   // Name of template param
	Provider        string                                                                                 `json:"provider,omitempty"`        // provider
	Range           *[]RequestItemConfigurationTemplatesImportsTheTemplatesProvidedV1TemplateParamsRange   `json:"range,omitempty"`           //
	Required        *bool                                                                                  `json:"required,omitempty"`        // Is param required
	Selection       *RequestItemConfigurationTemplatesImportsTheTemplatesProvidedV1TemplateParamsSelection `json:"selection,omitempty"`       //
}
type RequestItemConfigurationTemplatesImportsTheTemplatesProvidedV1TemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       // UUID of range
	MaxValue *int   `json:"maxValue,omitempty"` // Max value of range
	MinValue *int   `json:"minValue,omitempty"` // Min value of range
}
type RequestItemConfigurationTemplatesImportsTheTemplatesProvidedV1TemplateParamsSelection struct {
	DefaultSelectedValues []string                                                                                              `json:"defaultSelectedValues,omitempty"` // Default selection values
	ID                    string                                                                                                `json:"id,omitempty"`                    // UUID of selection
	SelectionType         string                                                                                                `json:"selectionType,omitempty"`         // Type of selection(SINGLE_SELECT or MULTI_SELECT)
	SelectionValues       *RequestItemConfigurationTemplatesImportsTheTemplatesProvidedV1TemplateParamsSelectionSelectionValues `json:"selectionValues,omitempty"`       // Selection values
}
type RequestItemConfigurationTemplatesImportsTheTemplatesProvidedV1TemplateParamsSelectionSelectionValues interface{}
type RequestItemConfigurationTemplatesImportsTheTemplatesProvidedV1ValidationErrors struct {
	RollbackTemplateErrors *RequestItemConfigurationTemplatesImportsTheTemplatesProvidedV1ValidationErrorsRollbackTemplateErrors `json:"rollbackTemplateErrors,omitempty"` // Validation or design conflicts errors of rollback template
	TemplateErrors         *RequestItemConfigurationTemplatesImportsTheTemplatesProvidedV1ValidationErrorsTemplateErrors         `json:"templateErrors,omitempty"`         // Validation or design conflicts errors
	TemplateID             string                                                                                                `json:"templateId,omitempty"`             // UUID of template
	TemplateVersion        string                                                                                                `json:"templateVersion,omitempty"`        // Current version of template
}
type RequestItemConfigurationTemplatesImportsTheTemplatesProvidedV1ValidationErrorsRollbackTemplateErrors interface{}
type RequestItemConfigurationTemplatesImportsTheTemplatesProvidedV1ValidationErrorsTemplateErrors interface{}
type RequestConfigurationTemplatesCreateTemplateV1 struct {
	Tags                    *[]RequestConfigurationTemplatesCreateTemplateV1Tags                   `json:"tags,omitempty"`                    //
	Author                  string                                                                 `json:"author,omitempty"`                  // Author of template
	Composite               *bool                                                                  `json:"composite,omitempty"`               // Is it composite template
	ContainingTemplates     *[]RequestConfigurationTemplatesCreateTemplateV1ContainingTemplates    `json:"containingTemplates,omitempty"`     //
	CreateTime              *int                                                                   `json:"createTime,omitempty"`              // Create time of template
	CustomParamsOrder       *bool                                                                  `json:"customParamsOrder,omitempty"`       // Custom Params Order
	Description             string                                                                 `json:"description,omitempty"`             // Description of template
	DeviceTypes             *[]RequestConfigurationTemplatesCreateTemplateV1DeviceTypes            `json:"deviceTypes,omitempty"`             //
	FailurePolicy           string                                                                 `json:"failurePolicy,omitempty"`           // Define failure policy if template provisioning fails
	ID                      string                                                                 `json:"id,omitempty"`                      // UUID of template
	Language                string                                                                 `json:"language,omitempty"`                // Template language (JINJA or VELOCITY)
	LastUpdateTime          *int                                                                   `json:"lastUpdateTime,omitempty"`          // Update time of template
	LatestVersionTime       *int                                                                   `json:"latestVersionTime,omitempty"`       // Latest versioned template time
	Name                    string                                                                 `json:"name,omitempty"`                    // Name of template
	ParentTemplateID        string                                                                 `json:"parentTemplateId,omitempty"`        // Parent templateID
	ProjectID               string                                                                 `json:"projectId,omitempty"`               // Project UUID
	ProjectName             string                                                                 `json:"projectName,omitempty"`             // Project name
	RollbackTemplateContent string                                                                 `json:"rollbackTemplateContent,omitempty"` // Rollback template content
	RollbackTemplateParams  *[]RequestConfigurationTemplatesCreateTemplateV1RollbackTemplateParams `json:"rollbackTemplateParams,omitempty"`  //
	SoftwareType            string                                                                 `json:"softwareType,omitempty"`            // Applicable device software type
	SoftwareVariant         string                                                                 `json:"softwareVariant,omitempty"`         // Applicable device software variant
	SoftwareVersion         string                                                                 `json:"softwareVersion,omitempty"`         // Applicable device software version
	TemplateContent         string                                                                 `json:"templateContent,omitempty"`         // Template content
	TemplateParams          *[]RequestConfigurationTemplatesCreateTemplateV1TemplateParams         `json:"templateParams,omitempty"`          //
	ValidationErrors        *RequestConfigurationTemplatesCreateTemplateV1ValidationErrors         `json:"validationErrors,omitempty"`        //
	Version                 string                                                                 `json:"version,omitempty"`                 // Current version of template
}
type RequestConfigurationTemplatesCreateTemplateV1Tags struct {
	ID   string `json:"id,omitempty"`   // UUID of tag
	Name string `json:"name,omitempty"` // Name of tag
}
type RequestConfigurationTemplatesCreateTemplateV1ContainingTemplates struct {
	Tags                   *[]RequestConfigurationTemplatesCreateTemplateV1ContainingTemplatesTags                   `json:"tags,omitempty"`                   //
	Composite              *bool                                                                                     `json:"composite,omitempty"`              // Is it composite template
	Description            string                                                                                    `json:"description,omitempty"`            // Description of template
	DeviceTypes            *[]RequestConfigurationTemplatesCreateTemplateV1ContainingTemplatesDeviceTypes            `json:"deviceTypes,omitempty"`            //
	ID                     string                                                                                    `json:"id,omitempty"`                     // UUID of template
	Language               string                                                                                    `json:"language,omitempty"`               // Template language (JINJA or VELOCITY)
	Name                   string                                                                                    `json:"name,omitempty"`                   // Name of template
	ProjectName            string                                                                                    `json:"projectName,omitempty"`            // Project name
	RollbackTemplateParams *[]RequestConfigurationTemplatesCreateTemplateV1ContainingTemplatesRollbackTemplateParams `json:"rollbackTemplateParams,omitempty"` //
	TemplateContent        string                                                                                    `json:"templateContent,omitempty"`        // Template content
	TemplateParams         *[]RequestConfigurationTemplatesCreateTemplateV1ContainingTemplatesTemplateParams         `json:"templateParams,omitempty"`         //
	Version                string                                                                                    `json:"version,omitempty"`                // Current version of template
}
type RequestConfigurationTemplatesCreateTemplateV1ContainingTemplatesTags struct {
	ID   string `json:"id,omitempty"`   // UUID of tag
	Name string `json:"name,omitempty"` // Name of tag
}
type RequestConfigurationTemplatesCreateTemplateV1ContainingTemplatesDeviceTypes struct {
	ProductFamily string `json:"productFamily,omitempty"` // Device family
	ProductSeries string `json:"productSeries,omitempty"` // Device series
	ProductType   string `json:"productType,omitempty"`   // Device type
}
type RequestConfigurationTemplatesCreateTemplateV1ContainingTemplatesRollbackTemplateParams struct {
	Binding         string                                                                                           `json:"binding,omitempty"`         // Bind to source
	CustomOrder     *int                                                                                             `json:"customOrder,omitempty"`     // CustomOrder of template param
	DataType        string                                                                                           `json:"dataType,omitempty"`        // Datatype of template param
	DefaultValue    string                                                                                           `json:"defaultValue,omitempty"`    // Default value of template param
	Description     string                                                                                           `json:"description,omitempty"`     // Description of template param
	DisplayName     string                                                                                           `json:"displayName,omitempty"`     // Display name of param
	Group           string                                                                                           `json:"group,omitempty"`           // group
	ID              string                                                                                           `json:"id,omitempty"`              // UUID of template param
	InstructionText string                                                                                           `json:"instructionText,omitempty"` // Instruction text for param
	Key             string                                                                                           `json:"key,omitempty"`             // key
	NotParam        *bool                                                                                            `json:"notParam,omitempty"`        // Is it not a variable
	Order           *int                                                                                             `json:"order,omitempty"`           // Order of template param
	ParamArray      *bool                                                                                            `json:"paramArray,omitempty"`      // Is it an array
	ParameterName   string                                                                                           `json:"parameterName,omitempty"`   // Name of template param
	Provider        string                                                                                           `json:"provider,omitempty"`        // provider
	Range           *[]RequestConfigurationTemplatesCreateTemplateV1ContainingTemplatesRollbackTemplateParamsRange   `json:"range,omitempty"`           //
	Required        *bool                                                                                            `json:"required,omitempty"`        // Is param required
	Selection       *RequestConfigurationTemplatesCreateTemplateV1ContainingTemplatesRollbackTemplateParamsSelection `json:"selection,omitempty"`       //
}
type RequestConfigurationTemplatesCreateTemplateV1ContainingTemplatesRollbackTemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       // UUID of range
	MaxValue *int   `json:"maxValue,omitempty"` // Max value of range
	MinValue *int   `json:"minValue,omitempty"` // Min value of range
}
type RequestConfigurationTemplatesCreateTemplateV1ContainingTemplatesRollbackTemplateParamsSelection struct {
	DefaultSelectedValues []string                                                                                                        `json:"defaultSelectedValues,omitempty"` // Default selection values
	ID                    string                                                                                                          `json:"id,omitempty"`                    // UUID of selection
	SelectionType         string                                                                                                          `json:"selectionType,omitempty"`         // Type of selection(SINGLE_SELECT or MULTI_SELECT)
	SelectionValues       *RequestConfigurationTemplatesCreateTemplateV1ContainingTemplatesRollbackTemplateParamsSelectionSelectionValues `json:"selectionValues,omitempty"`       // Selection values
}
type RequestConfigurationTemplatesCreateTemplateV1ContainingTemplatesRollbackTemplateParamsSelectionSelectionValues interface{}
type RequestConfigurationTemplatesCreateTemplateV1ContainingTemplatesTemplateParams struct {
	Binding         string                                                                                   `json:"binding,omitempty"`         // Bind to source
	CustomOrder     *int                                                                                     `json:"customOrder,omitempty"`     // CustomOrder of template param
	DataType        string                                                                                   `json:"dataType,omitempty"`        // Datatype of template param
	DefaultValue    string                                                                                   `json:"defaultValue,omitempty"`    // Default value of template param
	Description     string                                                                                   `json:"description,omitempty"`     // Description of template param
	DisplayName     string                                                                                   `json:"displayName,omitempty"`     // Display name of param
	Group           string                                                                                   `json:"group,omitempty"`           // group
	ID              string                                                                                   `json:"id,omitempty"`              // UUID of template param
	InstructionText string                                                                                   `json:"instructionText,omitempty"` // Instruction text for param
	Key             string                                                                                   `json:"key,omitempty"`             // key
	NotParam        *bool                                                                                    `json:"notParam,omitempty"`        // Is it not a variable
	Order           *int                                                                                     `json:"order,omitempty"`           // Order of template param
	ParamArray      *bool                                                                                    `json:"paramArray,omitempty"`      // Is it an array
	ParameterName   string                                                                                   `json:"parameterName,omitempty"`   // Name of template param
	Provider        string                                                                                   `json:"provider,omitempty"`        // provider
	Range           *[]RequestConfigurationTemplatesCreateTemplateV1ContainingTemplatesTemplateParamsRange   `json:"range,omitempty"`           //
	Required        *bool                                                                                    `json:"required,omitempty"`        // Is param required
	Selection       *RequestConfigurationTemplatesCreateTemplateV1ContainingTemplatesTemplateParamsSelection `json:"selection,omitempty"`       //
}
type RequestConfigurationTemplatesCreateTemplateV1ContainingTemplatesTemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       // UUID of range
	MaxValue *int   `json:"maxValue,omitempty"` // Max value of range
	MinValue *int   `json:"minValue,omitempty"` // Min value of range
}
type RequestConfigurationTemplatesCreateTemplateV1ContainingTemplatesTemplateParamsSelection struct {
	DefaultSelectedValues []string                                                                                                `json:"defaultSelectedValues,omitempty"` // Default selection values
	ID                    string                                                                                                  `json:"id,omitempty"`                    // UUID of selection
	SelectionType         string                                                                                                  `json:"selectionType,omitempty"`         // Type of selection(SINGLE_SELECT or MULTI_SELECT)
	SelectionValues       *RequestConfigurationTemplatesCreateTemplateV1ContainingTemplatesTemplateParamsSelectionSelectionValues `json:"selectionValues,omitempty"`       // Selection values
}
type RequestConfigurationTemplatesCreateTemplateV1ContainingTemplatesTemplateParamsSelectionSelectionValues interface{}
type RequestConfigurationTemplatesCreateTemplateV1DeviceTypes struct {
	ProductFamily string `json:"productFamily,omitempty"` // Device family
	ProductSeries string `json:"productSeries,omitempty"` // Device series
	ProductType   string `json:"productType,omitempty"`   // Device type
}
type RequestConfigurationTemplatesCreateTemplateV1RollbackTemplateParams struct {
	Binding         string                                                                        `json:"binding,omitempty"`         // Bind to source
	CustomOrder     *int                                                                          `json:"customOrder,omitempty"`     // CustomOrder of template param
	DataType        string                                                                        `json:"dataType,omitempty"`        // Datatype of template param
	DefaultValue    string                                                                        `json:"defaultValue,omitempty"`    // Default value of template param
	Description     string                                                                        `json:"description,omitempty"`     // Description of template param
	DisplayName     string                                                                        `json:"displayName,omitempty"`     // Display name of param
	Group           string                                                                        `json:"group,omitempty"`           // group
	ID              string                                                                        `json:"id,omitempty"`              // UUID of template param
	InstructionText string                                                                        `json:"instructionText,omitempty"` // Instruction text for param
	Key             string                                                                        `json:"key,omitempty"`             // key
	NotParam        *bool                                                                         `json:"notParam,omitempty"`        // Is it not a variable
	Order           *int                                                                          `json:"order,omitempty"`           // Order of template param
	ParamArray      *bool                                                                         `json:"paramArray,omitempty"`      // Is it an array
	ParameterName   string                                                                        `json:"parameterName,omitempty"`   // Name of template param
	Provider        string                                                                        `json:"provider,omitempty"`        // provider
	Range           *[]RequestConfigurationTemplatesCreateTemplateV1RollbackTemplateParamsRange   `json:"range,omitempty"`           //
	Required        *bool                                                                         `json:"required,omitempty"`        // Is param required
	Selection       *RequestConfigurationTemplatesCreateTemplateV1RollbackTemplateParamsSelection `json:"selection,omitempty"`       //
}
type RequestConfigurationTemplatesCreateTemplateV1RollbackTemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       // UUID of range
	MaxValue *int   `json:"maxValue,omitempty"` // Max value of range
	MinValue *int   `json:"minValue,omitempty"` // Min value of range
}
type RequestConfigurationTemplatesCreateTemplateV1RollbackTemplateParamsSelection struct {
	DefaultSelectedValues []string                                                                                     `json:"defaultSelectedValues,omitempty"` // Default selection values
	ID                    string                                                                                       `json:"id,omitempty"`                    // UUID of selection
	SelectionType         string                                                                                       `json:"selectionType,omitempty"`         // Type of selection(SINGLE_SELECT or MULTI_SELECT)
	SelectionValues       *RequestConfigurationTemplatesCreateTemplateV1RollbackTemplateParamsSelectionSelectionValues `json:"selectionValues,omitempty"`       // Selection values
}
type RequestConfigurationTemplatesCreateTemplateV1RollbackTemplateParamsSelectionSelectionValues interface{}
type RequestConfigurationTemplatesCreateTemplateV1TemplateParams struct {
	Binding         string                                                                `json:"binding,omitempty"`         // Bind to source
	CustomOrder     *int                                                                  `json:"customOrder,omitempty"`     // CustomOrder of template param
	DataType        string                                                                `json:"dataType,omitempty"`        // Datatype of template param
	DefaultValue    string                                                                `json:"defaultValue,omitempty"`    // Default value of template param
	Description     string                                                                `json:"description,omitempty"`     // Description of template param
	DisplayName     string                                                                `json:"displayName,omitempty"`     // Display name of param
	Group           string                                                                `json:"group,omitempty"`           // group
	ID              string                                                                `json:"id,omitempty"`              // UUID of template param
	InstructionText string                                                                `json:"instructionText,omitempty"` // Instruction text for param
	Key             string                                                                `json:"key,omitempty"`             // key
	NotParam        *bool                                                                 `json:"notParam,omitempty"`        // Is it not a variable
	Order           *int                                                                  `json:"order,omitempty"`           // Order of template param
	ParamArray      *bool                                                                 `json:"paramArray,omitempty"`      // Is it an array
	ParameterName   string                                                                `json:"parameterName,omitempty"`   // Name of template param
	Provider        string                                                                `json:"provider,omitempty"`        // provider
	Range           *[]RequestConfigurationTemplatesCreateTemplateV1TemplateParamsRange   `json:"range,omitempty"`           //
	Required        *bool                                                                 `json:"required,omitempty"`        // Is param required
	Selection       *RequestConfigurationTemplatesCreateTemplateV1TemplateParamsSelection `json:"selection,omitempty"`       //
}
type RequestConfigurationTemplatesCreateTemplateV1TemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       // UUID of range
	MaxValue *int   `json:"maxValue,omitempty"` // Max value of range
	MinValue *int   `json:"minValue,omitempty"` // Min value of range
}
type RequestConfigurationTemplatesCreateTemplateV1TemplateParamsSelection struct {
	DefaultSelectedValues []string                                                                             `json:"defaultSelectedValues,omitempty"` // Default selection values
	ID                    string                                                                               `json:"id,omitempty"`                    // UUID of selection
	SelectionType         string                                                                               `json:"selectionType,omitempty"`         // Type of selection(SINGLE_SELECT or MULTI_SELECT)
	SelectionValues       *RequestConfigurationTemplatesCreateTemplateV1TemplateParamsSelectionSelectionValues `json:"selectionValues,omitempty"`       // Selection values
}
type RequestConfigurationTemplatesCreateTemplateV1TemplateParamsSelectionSelectionValues interface{}
type RequestConfigurationTemplatesCreateTemplateV1ValidationErrors struct {
	RollbackTemplateErrors *RequestConfigurationTemplatesCreateTemplateV1ValidationErrorsRollbackTemplateErrors `json:"rollbackTemplateErrors,omitempty"` // Validation or design conflicts errors of rollback template
	TemplateErrors         *RequestConfigurationTemplatesCreateTemplateV1ValidationErrorsTemplateErrors         `json:"templateErrors,omitempty"`         // Validation or design conflicts errors
	TemplateID             string                                                                               `json:"templateId,omitempty"`             // UUID of template
	TemplateVersion        string                                                                               `json:"templateVersion,omitempty"`        // Current version of template
}
type RequestConfigurationTemplatesCreateTemplateV1ValidationErrorsRollbackTemplateErrors interface{}
type RequestConfigurationTemplatesCreateTemplateV1ValidationErrorsTemplateErrors interface{}
type RequestConfigurationTemplatesUpdateTemplateV1 struct {
	Tags                    *[]RequestConfigurationTemplatesUpdateTemplateV1Tags                   `json:"tags,omitempty"`                    //
	Author                  string                                                                 `json:"author,omitempty"`                  // Author of template
	Composite               *bool                                                                  `json:"composite,omitempty"`               // Is it composite template
	ContainingTemplates     *[]RequestConfigurationTemplatesUpdateTemplateV1ContainingTemplates    `json:"containingTemplates,omitempty"`     //
	CreateTime              *int                                                                   `json:"createTime,omitempty"`              // Create time of template
	CustomParamsOrder       *bool                                                                  `json:"customParamsOrder,omitempty"`       // Custom Params Order
	Description             string                                                                 `json:"description,omitempty"`             // Description of template
	DeviceTypes             *[]RequestConfigurationTemplatesUpdateTemplateV1DeviceTypes            `json:"deviceTypes,omitempty"`             //
	FailurePolicy           string                                                                 `json:"failurePolicy,omitempty"`           // Define failure policy if template provisioning fails
	ID                      string                                                                 `json:"id,omitempty"`                      // UUID of template
	Language                string                                                                 `json:"language,omitempty"`                // Template language (JINJA or VELOCITY)
	LastUpdateTime          *int                                                                   `json:"lastUpdateTime,omitempty"`          // Update time of template
	LatestVersionTime       *int                                                                   `json:"latestVersionTime,omitempty"`       // Latest versioned template time
	Name                    string                                                                 `json:"name,omitempty"`                    // Name of template
	ParentTemplateID        string                                                                 `json:"parentTemplateId,omitempty"`        // Parent templateID
	ProjectID               string                                                                 `json:"projectId,omitempty"`               // Project UUID
	ProjectName             string                                                                 `json:"projectName,omitempty"`             // Project name
	RollbackTemplateContent string                                                                 `json:"rollbackTemplateContent,omitempty"` // Rollback template content
	RollbackTemplateParams  *[]RequestConfigurationTemplatesUpdateTemplateV1RollbackTemplateParams `json:"rollbackTemplateParams,omitempty"`  //
	SoftwareType            string                                                                 `json:"softwareType,omitempty"`            // Applicable device software type
	SoftwareVariant         string                                                                 `json:"softwareVariant,omitempty"`         // Applicable device software variant
	SoftwareVersion         string                                                                 `json:"softwareVersion,omitempty"`         // Applicable device software version
	TemplateContent         string                                                                 `json:"templateContent,omitempty"`         // Template content
	TemplateParams          *[]RequestConfigurationTemplatesUpdateTemplateV1TemplateParams         `json:"templateParams,omitempty"`          //
	ValidationErrors        *RequestConfigurationTemplatesUpdateTemplateV1ValidationErrors         `json:"validationErrors,omitempty"`        //
	Version                 string                                                                 `json:"version,omitempty"`                 // Current version of template
}
type RequestConfigurationTemplatesUpdateTemplateV1Tags struct {
	ID   string `json:"id,omitempty"`   // UUID of tag
	Name string `json:"name,omitempty"` // Name of tag
}
type RequestConfigurationTemplatesUpdateTemplateV1ContainingTemplates struct {
	Tags                   *[]RequestConfigurationTemplatesUpdateTemplateV1ContainingTemplatesTags                   `json:"tags,omitempty"`                   //
	Composite              *bool                                                                                     `json:"composite,omitempty"`              // Is it composite template
	Description            string                                                                                    `json:"description,omitempty"`            // Description of template
	DeviceTypes            *[]RequestConfigurationTemplatesUpdateTemplateV1ContainingTemplatesDeviceTypes            `json:"deviceTypes,omitempty"`            //
	ID                     string                                                                                    `json:"id,omitempty"`                     // UUID of template
	Language               string                                                                                    `json:"language,omitempty"`               // Template language (JINJA or VELOCITY)
	Name                   string                                                                                    `json:"name,omitempty"`                   // Name of template
	ProjectName            string                                                                                    `json:"projectName,omitempty"`            // Project name
	RollbackTemplateParams *[]RequestConfigurationTemplatesUpdateTemplateV1ContainingTemplatesRollbackTemplateParams `json:"rollbackTemplateParams,omitempty"` //
	TemplateContent        string                                                                                    `json:"templateContent,omitempty"`        // Template content
	TemplateParams         *[]RequestConfigurationTemplatesUpdateTemplateV1ContainingTemplatesTemplateParams         `json:"templateParams,omitempty"`         //
	Version                string                                                                                    `json:"version,omitempty"`                // Current version of template
}
type RequestConfigurationTemplatesUpdateTemplateV1ContainingTemplatesTags struct {
	ID   string `json:"id,omitempty"`   // UUID of tag
	Name string `json:"name,omitempty"` // Name of tag
}
type RequestConfigurationTemplatesUpdateTemplateV1ContainingTemplatesDeviceTypes struct {
	ProductFamily string `json:"productFamily,omitempty"` // Device family
	ProductSeries string `json:"productSeries,omitempty"` // Device series
	ProductType   string `json:"productType,omitempty"`   // Device type
}
type RequestConfigurationTemplatesUpdateTemplateV1ContainingTemplatesRollbackTemplateParams struct {
	Binding         string                                                                                           `json:"binding,omitempty"`         // Bind to source
	CustomOrder     *int                                                                                             `json:"customOrder,omitempty"`     // CustomOrder of template param
	DataType        string                                                                                           `json:"dataType,omitempty"`        // Datatype of template param
	DefaultValue    string                                                                                           `json:"defaultValue,omitempty"`    // Default value of template param
	Description     string                                                                                           `json:"description,omitempty"`     // Description of template param
	DisplayName     string                                                                                           `json:"displayName,omitempty"`     // Display name of param
	Group           string                                                                                           `json:"group,omitempty"`           // group
	ID              string                                                                                           `json:"id,omitempty"`              // UUID of template param
	InstructionText string                                                                                           `json:"instructionText,omitempty"` // Instruction text for param
	Key             string                                                                                           `json:"key,omitempty"`             // key
	NotParam        *bool                                                                                            `json:"notParam,omitempty"`        // Is it not a variable
	Order           *int                                                                                             `json:"order,omitempty"`           // Order of template param
	ParamArray      *bool                                                                                            `json:"paramArray,omitempty"`      // Is it an array
	ParameterName   string                                                                                           `json:"parameterName,omitempty"`   // Name of template param
	Provider        string                                                                                           `json:"provider,omitempty"`        // provider
	Range           *[]RequestConfigurationTemplatesUpdateTemplateV1ContainingTemplatesRollbackTemplateParamsRange   `json:"range,omitempty"`           //
	Required        *bool                                                                                            `json:"required,omitempty"`        // Is param required
	Selection       *RequestConfigurationTemplatesUpdateTemplateV1ContainingTemplatesRollbackTemplateParamsSelection `json:"selection,omitempty"`       //
}
type RequestConfigurationTemplatesUpdateTemplateV1ContainingTemplatesRollbackTemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       // UUID of range
	MaxValue *int   `json:"maxValue,omitempty"` // Max value of range
	MinValue *int   `json:"minValue,omitempty"` // Min value of range
}
type RequestConfigurationTemplatesUpdateTemplateV1ContainingTemplatesRollbackTemplateParamsSelection struct {
	DefaultSelectedValues []string                                                                                                        `json:"defaultSelectedValues,omitempty"` // Default selection values
	ID                    string                                                                                                          `json:"id,omitempty"`                    // UUID of selection
	SelectionType         string                                                                                                          `json:"selectionType,omitempty"`         // Type of selection(SINGLE_SELECT or MULTI_SELECT)
	SelectionValues       *RequestConfigurationTemplatesUpdateTemplateV1ContainingTemplatesRollbackTemplateParamsSelectionSelectionValues `json:"selectionValues,omitempty"`       // Selection values
}
type RequestConfigurationTemplatesUpdateTemplateV1ContainingTemplatesRollbackTemplateParamsSelectionSelectionValues interface{}
type RequestConfigurationTemplatesUpdateTemplateV1ContainingTemplatesTemplateParams struct {
	Binding         string                                                                                   `json:"binding,omitempty"`         // Bind to source
	CustomOrder     *int                                                                                     `json:"customOrder,omitempty"`     // CustomOrder of template param
	DataType        string                                                                                   `json:"dataType,omitempty"`        // Datatype of template param
	DefaultValue    string                                                                                   `json:"defaultValue,omitempty"`    // Default value of template param
	Description     string                                                                                   `json:"description,omitempty"`     // Description of template param
	DisplayName     string                                                                                   `json:"displayName,omitempty"`     // Display name of param
	Group           string                                                                                   `json:"group,omitempty"`           // group
	ID              string                                                                                   `json:"id,omitempty"`              // UUID of template param
	InstructionText string                                                                                   `json:"instructionText,omitempty"` // Instruction text for param
	Key             string                                                                                   `json:"key,omitempty"`             // key
	NotParam        *bool                                                                                    `json:"notParam,omitempty"`        // Is it not a variable
	Order           *int                                                                                     `json:"order,omitempty"`           // Order of template param
	ParamArray      *bool                                                                                    `json:"paramArray,omitempty"`      // Is it an array
	ParameterName   string                                                                                   `json:"parameterName,omitempty"`   // Name of template param
	Provider        string                                                                                   `json:"provider,omitempty"`        // provider
	Range           *[]RequestConfigurationTemplatesUpdateTemplateV1ContainingTemplatesTemplateParamsRange   `json:"range,omitempty"`           //
	Required        *bool                                                                                    `json:"required,omitempty"`        // Is param required
	Selection       *RequestConfigurationTemplatesUpdateTemplateV1ContainingTemplatesTemplateParamsSelection `json:"selection,omitempty"`       //
}
type RequestConfigurationTemplatesUpdateTemplateV1ContainingTemplatesTemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       // UUID of range
	MaxValue *int   `json:"maxValue,omitempty"` // Max value of range
	MinValue *int   `json:"minValue,omitempty"` // Min value of range
}
type RequestConfigurationTemplatesUpdateTemplateV1ContainingTemplatesTemplateParamsSelection struct {
	DefaultSelectedValues []string                                                                                                `json:"defaultSelectedValues,omitempty"` // Default selection values
	ID                    string                                                                                                  `json:"id,omitempty"`                    // UUID of selection
	SelectionType         string                                                                                                  `json:"selectionType,omitempty"`         // Type of selection(SINGLE_SELECT or MULTI_SELECT)
	SelectionValues       *RequestConfigurationTemplatesUpdateTemplateV1ContainingTemplatesTemplateParamsSelectionSelectionValues `json:"selectionValues,omitempty"`       // Selection values
}
type RequestConfigurationTemplatesUpdateTemplateV1ContainingTemplatesTemplateParamsSelectionSelectionValues interface{}
type RequestConfigurationTemplatesUpdateTemplateV1DeviceTypes struct {
	ProductFamily string `json:"productFamily,omitempty"` // Device family
	ProductSeries string `json:"productSeries,omitempty"` // Device series
	ProductType   string `json:"productType,omitempty"`   // Device type
}
type RequestConfigurationTemplatesUpdateTemplateV1RollbackTemplateParams struct {
	Binding         string                                                                        `json:"binding,omitempty"`         // Bind to source
	CustomOrder     *int                                                                          `json:"customOrder,omitempty"`     // CustomOrder of template param
	DataType        string                                                                        `json:"dataType,omitempty"`        // Datatype of template param
	DefaultValue    string                                                                        `json:"defaultValue,omitempty"`    // Default value of template param
	Description     string                                                                        `json:"description,omitempty"`     // Description of template param
	DisplayName     string                                                                        `json:"displayName,omitempty"`     // Display name of param
	Group           string                                                                        `json:"group,omitempty"`           // group
	ID              string                                                                        `json:"id,omitempty"`              // UUID of template param
	InstructionText string                                                                        `json:"instructionText,omitempty"` // Instruction text for param
	Key             string                                                                        `json:"key,omitempty"`             // key
	NotParam        *bool                                                                         `json:"notParam,omitempty"`        // Is it not a variable
	Order           *int                                                                          `json:"order,omitempty"`           // Order of template param
	ParamArray      *bool                                                                         `json:"paramArray,omitempty"`      // Is it an array
	ParameterName   string                                                                        `json:"parameterName,omitempty"`   // Name of template param
	Provider        string                                                                        `json:"provider,omitempty"`        // provider
	Range           *[]RequestConfigurationTemplatesUpdateTemplateV1RollbackTemplateParamsRange   `json:"range,omitempty"`           //
	Required        *bool                                                                         `json:"required,omitempty"`        // Is param required
	Selection       *RequestConfigurationTemplatesUpdateTemplateV1RollbackTemplateParamsSelection `json:"selection,omitempty"`       //
}
type RequestConfigurationTemplatesUpdateTemplateV1RollbackTemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       // UUID of range
	MaxValue *int   `json:"maxValue,omitempty"` // Max value of range
	MinValue *int   `json:"minValue,omitempty"` // Min value of range
}
type RequestConfigurationTemplatesUpdateTemplateV1RollbackTemplateParamsSelection struct {
	DefaultSelectedValues []string                                                                                     `json:"defaultSelectedValues,omitempty"` // Default selection values
	ID                    string                                                                                       `json:"id,omitempty"`                    // UUID of selection
	SelectionType         string                                                                                       `json:"selectionType,omitempty"`         // Type of selection(SINGLE_SELECT or MULTI_SELECT)
	SelectionValues       *RequestConfigurationTemplatesUpdateTemplateV1RollbackTemplateParamsSelectionSelectionValues `json:"selectionValues,omitempty"`       // Selection values
}
type RequestConfigurationTemplatesUpdateTemplateV1RollbackTemplateParamsSelectionSelectionValues interface{}
type RequestConfigurationTemplatesUpdateTemplateV1TemplateParams struct {
	Binding         string                                                                `json:"binding,omitempty"`         // Bind to source
	CustomOrder     *int                                                                  `json:"customOrder,omitempty"`     // CustomOrder of template param
	DataType        string                                                                `json:"dataType,omitempty"`        // Datatype of template param
	DefaultValue    string                                                                `json:"defaultValue,omitempty"`    // Default value of template param
	Description     string                                                                `json:"description,omitempty"`     // Description of template param
	DisplayName     string                                                                `json:"displayName,omitempty"`     // Display name of param
	Group           string                                                                `json:"group,omitempty"`           // group
	ID              string                                                                `json:"id,omitempty"`              // UUID of template param
	InstructionText string                                                                `json:"instructionText,omitempty"` // Instruction text for param
	Key             string                                                                `json:"key,omitempty"`             // key
	NotParam        *bool                                                                 `json:"notParam,omitempty"`        // Is it not a variable
	Order           *int                                                                  `json:"order,omitempty"`           // Order of template param
	ParamArray      *bool                                                                 `json:"paramArray,omitempty"`      // Is it an array
	ParameterName   string                                                                `json:"parameterName,omitempty"`   // Name of template param
	Provider        string                                                                `json:"provider,omitempty"`        // provider
	Range           *[]RequestConfigurationTemplatesUpdateTemplateV1TemplateParamsRange   `json:"range,omitempty"`           //
	Required        *bool                                                                 `json:"required,omitempty"`        // Is param required
	Selection       *RequestConfigurationTemplatesUpdateTemplateV1TemplateParamsSelection `json:"selection,omitempty"`       //
}
type RequestConfigurationTemplatesUpdateTemplateV1TemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       // UUID of range
	MaxValue *int   `json:"maxValue,omitempty"` // Max value of range
	MinValue *int   `json:"minValue,omitempty"` // Min value of range
}
type RequestConfigurationTemplatesUpdateTemplateV1TemplateParamsSelection struct {
	DefaultSelectedValues []string                                                                             `json:"defaultSelectedValues,omitempty"` // Default selection values
	ID                    string                                                                               `json:"id,omitempty"`                    // UUID of selection
	SelectionType         string                                                                               `json:"selectionType,omitempty"`         // Type of selection(SINGLE_SELECT or MULTI_SELECT)
	SelectionValues       *RequestConfigurationTemplatesUpdateTemplateV1TemplateParamsSelectionSelectionValues `json:"selectionValues,omitempty"`       // Selection values
}
type RequestConfigurationTemplatesUpdateTemplateV1TemplateParamsSelectionSelectionValues interface{}
type RequestConfigurationTemplatesUpdateTemplateV1ValidationErrors struct {
	RollbackTemplateErrors *RequestConfigurationTemplatesUpdateTemplateV1ValidationErrorsRollbackTemplateErrors `json:"rollbackTemplateErrors,omitempty"` // Validation or design conflicts errors of rollback template
	TemplateErrors         *RequestConfigurationTemplatesUpdateTemplateV1ValidationErrorsTemplateErrors         `json:"templateErrors,omitempty"`         // Validation or design conflicts errors
	TemplateID             string                                                                               `json:"templateId,omitempty"`             // UUID of template
	TemplateVersion        string                                                                               `json:"templateVersion,omitempty"`        // Current version of template
}
type RequestConfigurationTemplatesUpdateTemplateV1ValidationErrorsRollbackTemplateErrors interface{}
type RequestConfigurationTemplatesUpdateTemplateV1ValidationErrorsTemplateErrors interface{}
type RequestConfigurationTemplatesDeployTemplateV1 struct {
	ForcePushTemplate            *bool                                                      `json:"forcePushTemplate,omitempty"`            //
	IsComposite                  *bool                                                      `json:"isComposite,omitempty"`                  // Composite template flag
	MainTemplateID               string                                                     `json:"mainTemplateId,omitempty"`               // Main template UUID of versioned template
	MemberTemplateDeploymentInfo []string                                                   `json:"memberTemplateDeploymentInfo,omitempty"` // memberTemplateDeploymentInfo
	TargetInfo                   *[]RequestConfigurationTemplatesDeployTemplateV1TargetInfo `json:"targetInfo,omitempty"`                   //
	TemplateID                   string                                                     `json:"templateId,omitempty"`                   // UUID of template to be provisioned
}
type RequestConfigurationTemplatesDeployTemplateV1TargetInfo struct {
	HostName            string                                                         `json:"hostName,omitempty"`            // Hostname of device is required if targetType is MANAGED_DEVICE_HOSTNAME
	ID                  string                                                         `json:"id,omitempty"`                  // UUID of target is required if targetType is MANAGED_DEVICE_UUID
	Params              *RequestConfigurationTemplatesDeployTemplateV1TargetInfoParams `json:"params,omitempty"`              // Template params/values to be provisioned
	ResourceParams      []string                                                       `json:"resourceParams,omitempty"`      // Resource params to be provisioned. Refer to features page for usage details
	Type                string                                                         `json:"type,omitempty"`                // Target type of device
	VersionedTemplateID string                                                         `json:"versionedTemplateId,omitempty"` // Versioned templateUUID to be provisioned
}
type RequestConfigurationTemplatesDeployTemplateV1TargetInfoParams interface{}
type RequestConfigurationTemplatesExportsTheTemplatesForAGivenCriteriaV1 []RequestItemConfigurationTemplatesExportsTheTemplatesForAGivenCriteriaV1 // Array of RequestConfigurationTemplatesExportsTheTemplatesForAGivenCriteriaV1
type RequestItemConfigurationTemplatesExportsTheTemplatesForAGivenCriteriaV1 interface{}
type RequestConfigurationTemplatesPreviewTemplateV1 struct {
	DeviceID       string                                                        `json:"deviceId,omitempty"`       // UUID of device to get template preview
	Params         *RequestConfigurationTemplatesPreviewTemplateV1Params         `json:"params,omitempty"`         // Params to render preview
	ResourceParams *RequestConfigurationTemplatesPreviewTemplateV1ResourceParams `json:"resourceParams,omitempty"` // Resource params to render preview
	TemplateID     string                                                        `json:"templateId,omitempty"`     // UUID of template to get template preview
}
type RequestConfigurationTemplatesPreviewTemplateV1Params interface{}
type RequestConfigurationTemplatesPreviewTemplateV1ResourceParams interface{}
type RequestConfigurationTemplatesVersionTemplateV1 struct {
	Comments   string `json:"comments,omitempty"`   // Template version comments
	TemplateID string `json:"templateId,omitempty"` // UUID of template
}
type RequestConfigurationTemplatesDeployTemplateV2 struct {
	ForcePushTemplate            *bool                                                      `json:"forcePushTemplate,omitempty"`            //
	IsComposite                  *bool                                                      `json:"isComposite,omitempty"`                  // Composite template flag
	MainTemplateID               string                                                     `json:"mainTemplateId,omitempty"`               // Main template UUID of versioned template
	MemberTemplateDeploymentInfo []string                                                   `json:"memberTemplateDeploymentInfo,omitempty"` // memberTemplateDeploymentInfo
	TargetInfo                   *[]RequestConfigurationTemplatesDeployTemplateV2TargetInfo `json:"targetInfo,omitempty"`                   //
	TemplateID                   string                                                     `json:"templateId,omitempty"`                   // UUID of template to be provisioned
}
type RequestConfigurationTemplatesDeployTemplateV2TargetInfo struct {
	HostName            string                                                         `json:"hostName,omitempty"`            // Hostname of device is required if targetType is MANAGED_DEVICE_HOSTNAME
	ID                  string                                                         `json:"id,omitempty"`                  // UUID of target is required if targetType is MANAGED_DEVICE_UUID
	Params              *RequestConfigurationTemplatesDeployTemplateV2TargetInfoParams `json:"params,omitempty"`              // Template params/values to be provisioned
	ResourceParams      []string                                                       `json:"resourceParams,omitempty"`      // Resource params to be provisioned. Refer to features page for usage details
	Type                string                                                         `json:"type,omitempty"`                // Target type of device
	VersionedTemplateID string                                                         `json:"versionedTemplateId,omitempty"` // Versioned templateUUID to be provisioned
}
type RequestConfigurationTemplatesDeployTemplateV2TargetInfoParams map[string]interface{}

//GetsAListOfProjectsV1 Gets a list of projects - 4f80-08c2-400b-98ee
/* List the projects


@param GetsAListOfProjectsV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!gets-a-list-of-projects-v1
*/
func (s *ConfigurationTemplatesService) GetsAListOfProjectsV1(GetsAListOfProjectsV1QueryParams *GetsAListOfProjectsV1QueryParams) (*ResponseConfigurationTemplatesGetsAListOfProjectsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/template-programmer/project"

	queryString, _ := query.Values(GetsAListOfProjectsV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseConfigurationTemplatesGetsAListOfProjectsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetsAListOfProjectsV1(GetsAListOfProjectsV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetsAListOfProjectsV1")
	}

	result := response.Result().(*ResponseConfigurationTemplatesGetsAListOfProjectsV1)
	return result, response, err

}

//GetsTheDetailsOfAGivenProjectV1 Gets the details of a given project. - dd91-a8c0-436a-82d9
/* Get the details of the given project by its id.


@param projectID projectId path parameter. projectId(UUID) of project to get project details


Documentation Link: https://developer.cisco.com/docs/dna-center/#!gets-the-details-of-a-given-project-v1
*/
func (s *ConfigurationTemplatesService) GetsTheDetailsOfAGivenProjectV1(projectID string) (*ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/template-programmer/project/{projectId}"
	path = strings.Replace(path, "{projectId}", fmt.Sprintf("%v", projectID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetsTheDetailsOfAGivenProjectV1(projectID)
		}
		return nil, response, fmt.Errorf("error with operation GetsTheDetailsOfAGivenProjectV1")
	}

	result := response.Result().(*ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectV1)
	return result, response, err

}

//GetsTheTemplatesAvailableV1 Gets the templates available - e286-e848-47bb-a77e
/* List the templates available


@param GetsTheTemplatesAvailableV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!gets-the-templates-available-v1
*/
func (s *ConfigurationTemplatesService) GetsTheTemplatesAvailableV1(GetsTheTemplatesAvailableV1QueryParams *GetsTheTemplatesAvailableV1QueryParams) (*ResponseConfigurationTemplatesGetsTheTemplatesAvailableV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/template-programmer/template"

	queryString, _ := query.Values(GetsTheTemplatesAvailableV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseConfigurationTemplatesGetsTheTemplatesAvailableV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetsTheTemplatesAvailableV1(GetsTheTemplatesAvailableV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetsTheTemplatesAvailableV1")
	}

	result := response.Result().(*ResponseConfigurationTemplatesGetsTheTemplatesAvailableV1)
	return result, response, err

}

//StatusOfTemplateDeploymentV1 Status of template deployment - 078e-f800-49b8-80f1
/* API to retrieve the status of template deployment.


@param deploymentID deploymentId path parameter. UUID of deployment to retrieve template deployment status


Documentation Link: https://developer.cisco.com/docs/dna-center/#!status-of-template-deployment-v1
*/
func (s *ConfigurationTemplatesService) StatusOfTemplateDeploymentV1(deploymentID string) (*ResponseConfigurationTemplatesStatusOfTemplateDeploymentV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/template-programmer/template/deploy/status/{deploymentId}"
	path = strings.Replace(path, "{deploymentId}", fmt.Sprintf("%v", deploymentID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseConfigurationTemplatesStatusOfTemplateDeploymentV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.StatusOfTemplateDeploymentV1(deploymentID)
		}
		return nil, response, fmt.Errorf("error with operation StatusOfTemplateDeploymentV1")
	}

	result := response.Result().(*ResponseConfigurationTemplatesStatusOfTemplateDeploymentV1)
	return result, response, err

}

//GetsAllTheVersionsOfAGivenTemplateV1 Gets all the versions of a given template - 0191-2b65-45fb-8891
/* Get all the versions of template by its id


@param templateID templateId path parameter. templateId(UUID) to get list of versioned templates


Documentation Link: https://developer.cisco.com/docs/dna-center/#!gets-all-the-versions-of-a-given-template-v1
*/
func (s *ConfigurationTemplatesService) GetsAllTheVersionsOfAGivenTemplateV1(templateID string) (*ResponseConfigurationTemplatesGetsAllTheVersionsOfAGivenTemplateV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/template-programmer/template/version/{templateId}"
	path = strings.Replace(path, "{templateId}", fmt.Sprintf("%v", templateID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseConfigurationTemplatesGetsAllTheVersionsOfAGivenTemplateV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetsAllTheVersionsOfAGivenTemplateV1(templateID)
		}
		return nil, response, fmt.Errorf("error with operation GetsAllTheVersionsOfAGivenTemplateV1")
	}

	result := response.Result().(*ResponseConfigurationTemplatesGetsAllTheVersionsOfAGivenTemplateV1)
	return result, response, err

}

//GetsDetailsOfAGivenTemplateV1 Gets details of a given template - 8c82-5900-40d9-8137
/* Details of the template by its id


@param templateID templateId path parameter. TemplateId(UUID) to get details of the template

@param GetsDetailsOfAGivenTemplateV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!gets-details-of-a-given-template-v1
*/
func (s *ConfigurationTemplatesService) GetsDetailsOfAGivenTemplateV1(templateID string, GetsDetailsOfAGivenTemplateV1QueryParams *GetsDetailsOfAGivenTemplateV1QueryParams) (*ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/template-programmer/template/{templateId}"
	path = strings.Replace(path, "{templateId}", fmt.Sprintf("%v", templateID), -1)

	queryString, _ := query.Values(GetsDetailsOfAGivenTemplateV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetsDetailsOfAGivenTemplateV1(templateID, GetsDetailsOfAGivenTemplateV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetsDetailsOfAGivenTemplateV1")
	}

	result := response.Result().(*ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateV1)
	return result, response, err

}

//GetProjectsDetailsV2 Get project(s) details - 9a8c-aa6d-459b-a4a2
/* Get project(s) details


@param GetProjectsDetailsV2QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-projects-details-v2
*/
func (s *ConfigurationTemplatesService) GetProjectsDetailsV2(GetProjectsDetailsV2QueryParams *GetProjectsDetailsV2QueryParams) (*ResponseConfigurationTemplatesGetProjectsDetailsV2, *resty.Response, error) {
	path := "/dna/intent/api/v2/template-programmer/project"

	queryString, _ := query.Values(GetProjectsDetailsV2QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseConfigurationTemplatesGetProjectsDetailsV2{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetProjectsDetailsV2(GetProjectsDetailsV2QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetProjectsDetailsV2")
	}

	result := response.Result().(*ResponseConfigurationTemplatesGetProjectsDetailsV2)
	return result, response, err

}

//GetTemplatesDetailsV2 Get template(s) details - b0b6-ba49-43c8-9f45
/* Get template(s) details


@param GetTemplatesDetailsV2QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-templates-details-v2
*/
func (s *ConfigurationTemplatesService) GetTemplatesDetailsV2(GetTemplatesDetailsV2QueryParams *GetTemplatesDetailsV2QueryParams) (*ResponseConfigurationTemplatesGetTemplatesDetailsV2, *resty.Response, error) {
	path := "/dna/intent/api/v2/template-programmer/template"

	queryString, _ := query.Values(GetTemplatesDetailsV2QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseConfigurationTemplatesGetTemplatesDetailsV2{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTemplatesDetailsV2(GetTemplatesDetailsV2QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetTemplatesDetailsV2")
	}

	result := response.Result().(*ResponseConfigurationTemplatesGetTemplatesDetailsV2)
	return result, response, err

}

//CreatesACloneOfTheGivenTemplateV1 Creates a clone of the given template - 0384-4a0a-4ee8-bfc2
/* API to clone template


@param name name path parameter. Template name to clone template(Name should be different than existing template name within same project)

@param templateID templateId path parameter. UUID of the template to clone it

@param projectID projectId path parameter.
@param CreatesACloneOfTheGivenTemplateV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!creates-a-clone-of-the-given-template-v1
*/
func (s *ConfigurationTemplatesService) CreatesACloneOfTheGivenTemplateV1(name string, templateID string, projectID string, CreatesACloneOfTheGivenTemplateV1QueryParams *CreatesACloneOfTheGivenTemplateV1QueryParams) (*ResponseConfigurationTemplatesCreatesACloneOfTheGivenTemplateV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/template-programmer/clone/name/{name}/project/{projectId}/template/{templateId}"
	path = strings.Replace(path, "{name}", fmt.Sprintf("%v", name), -1)
	path = strings.Replace(path, "{templateId}", fmt.Sprintf("%v", templateID), -1)
	path = strings.Replace(path, "{projectId}", fmt.Sprintf("%v", projectID), -1)

	queryString, _ := query.Values(CreatesACloneOfTheGivenTemplateV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).
		SetResult(&ResponseConfigurationTemplatesCreatesACloneOfTheGivenTemplateV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreatesACloneOfTheGivenTemplateV1(name, templateID, projectID, CreatesACloneOfTheGivenTemplateV1QueryParams)
		}

		return nil, response, fmt.Errorf("error with operation CreatesACloneOfTheGivenTemplateV1")
	}

	result := response.Result().(*ResponseConfigurationTemplatesCreatesACloneOfTheGivenTemplateV1)
	return result, response, err

}

//CreateProjectV1 Create Project - 5788-a8a8-4aa9-b97a
/* This API is used to create a new project.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-project-v1
*/
func (s *ConfigurationTemplatesService) CreateProjectV1(requestConfigurationTemplatesCreateProjectV1 *RequestConfigurationTemplatesCreateProjectV1) (*ResponseConfigurationTemplatesCreateProjectV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/template-programmer/project"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestConfigurationTemplatesCreateProjectV1).
		SetResult(&ResponseConfigurationTemplatesCreateProjectV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateProjectV1(requestConfigurationTemplatesCreateProjectV1)
		}

		return nil, response, fmt.Errorf("error with operation CreateProjectV1")
	}

	result := response.Result().(*ResponseConfigurationTemplatesCreateProjectV1)
	return result, response, err

}

//ImportsTheProjectsProvidedV1 Imports the Projects provided - f59f-a8ad-42d9-8b4f
/* Imports the Projects provided in the DTO


@param ImportsTheProjectsProvidedV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!imports-the-projects-provided-v1
*/
func (s *ConfigurationTemplatesService) ImportsTheProjectsProvidedV1(ImportsTheProjectsProvidedV1QueryParams *ImportsTheProjectsProvidedV1QueryParams) (*ResponseConfigurationTemplatesImportsTheProjectsProvidedV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/template-programmer/project/importprojects"

	queryString, _ := query.Values(ImportsTheProjectsProvidedV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).
		SetResult(&ResponseConfigurationTemplatesImportsTheProjectsProvidedV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.ImportsTheProjectsProvidedV1(ImportsTheProjectsProvidedV1QueryParams)
		}

		return nil, response, fmt.Errorf("error with operation ImportsTheProjectsProvidedV1")
	}

	result := response.Result().(*ResponseConfigurationTemplatesImportsTheProjectsProvidedV1)
	return result, response, err

}

//ExportsTheProjectsForAGivenCriteriaV1 Exports the projects for a given criteria. - 67bc-e964-45f8-b720
/* Exports the projects for given projectNames.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!exports-the-projects-for-a-given-criteria-v1
*/
func (s *ConfigurationTemplatesService) ExportsTheProjectsForAGivenCriteriaV1(requestConfigurationTemplatesExportsTheProjectsForAGivenCriteriaV1 *RequestConfigurationTemplatesExportsTheProjectsForAGivenCriteriaV1) (*ResponseConfigurationTemplatesExportsTheProjectsForAGivenCriteriaV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/template-programmer/project/name/exportprojects"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestConfigurationTemplatesExportsTheProjectsForAGivenCriteriaV1).
		SetResult(&ResponseConfigurationTemplatesExportsTheProjectsForAGivenCriteriaV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.ExportsTheProjectsForAGivenCriteriaV1(requestConfigurationTemplatesExportsTheProjectsForAGivenCriteriaV1)
		}

		return nil, response, fmt.Errorf("error with operation ExportsTheProjectsForAGivenCriteriaV1")
	}

	result := response.Result().(*ResponseConfigurationTemplatesExportsTheProjectsForAGivenCriteriaV1)
	return result, response, err

}

//ImportsTheTemplatesProvidedV1 Imports the templates provided - 4d86-f92a-4a7b-90bb
/* Imports the templates provided in the DTO by project Name


@param projectName projectName path parameter. Project name to create template under the project

@param ImportsTheTemplatesProvidedV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!imports-the-templates-provided-v1
*/
func (s *ConfigurationTemplatesService) ImportsTheTemplatesProvidedV1(projectName string, requestConfigurationTemplatesImportsTheTemplatesProvidedV1 *RequestConfigurationTemplatesImportsTheTemplatesProvidedV1, ImportsTheTemplatesProvidedV1QueryParams *ImportsTheTemplatesProvidedV1QueryParams) (*ResponseConfigurationTemplatesImportsTheTemplatesProvidedV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/template-programmer/project/name/{projectName}/template/importtemplates"
	path = strings.Replace(path, "{projectName}", fmt.Sprintf("%v", projectName), -1)

	queryString, _ := query.Values(ImportsTheTemplatesProvidedV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetBody(requestConfigurationTemplatesImportsTheTemplatesProvidedV1).
		SetResult(&ResponseConfigurationTemplatesImportsTheTemplatesProvidedV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.ImportsTheTemplatesProvidedV1(projectName, requestConfigurationTemplatesImportsTheTemplatesProvidedV1, ImportsTheTemplatesProvidedV1QueryParams)
		}

		return nil, response, fmt.Errorf("error with operation ImportsTheTemplatesProvidedV1")
	}

	result := response.Result().(*ResponseConfigurationTemplatesImportsTheTemplatesProvidedV1)
	return result, response, err

}

//CreateTemplateV1 Create Template - ab94-a88b-4b0b-8d3d
/* API to create a template by project id.


@param projectID projectId path parameter. UUID of the project in which the template needs to be created


Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-template-v1
*/
func (s *ConfigurationTemplatesService) CreateTemplateV1(projectID string, requestConfigurationTemplatesCreateTemplateV1 *RequestConfigurationTemplatesCreateTemplateV1) (*ResponseConfigurationTemplatesCreateTemplateV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/template-programmer/project/{projectId}/template"
	path = strings.Replace(path, "{projectId}", fmt.Sprintf("%v", projectID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestConfigurationTemplatesCreateTemplateV1).
		SetResult(&ResponseConfigurationTemplatesCreateTemplateV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateTemplateV1(projectID, requestConfigurationTemplatesCreateTemplateV1)
		}

		return nil, response, fmt.Errorf("error with operation CreateTemplateV1")
	}

	result := response.Result().(*ResponseConfigurationTemplatesCreateTemplateV1)
	return result, response, err

}

//DeployTemplateV1 Deploy Template - 179f-09d8-430b-bee0
/* API to deploy a template.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!deploy-template-v1
*/
func (s *ConfigurationTemplatesService) DeployTemplateV1(requestConfigurationTemplatesDeployTemplateV1 *RequestConfigurationTemplatesDeployTemplateV1) (*ResponseConfigurationTemplatesDeployTemplateV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/template-programmer/template/deploy"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestConfigurationTemplatesDeployTemplateV1).
		SetResult(&ResponseConfigurationTemplatesDeployTemplateV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeployTemplateV1(requestConfigurationTemplatesDeployTemplateV1)
		}

		return nil, response, fmt.Errorf("error with operation DeployTemplateV1")
	}

	result := response.Result().(*ResponseConfigurationTemplatesDeployTemplateV1)
	return result, response, err

}

//ExportsTheTemplatesForAGivenCriteriaV1 Exports the templates for a given criteria. - a3a9-498f-4f48-a3c7
/* Exports the templates for given templateIds.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!exports-the-templates-for-a-given-criteria-v1
*/
func (s *ConfigurationTemplatesService) ExportsTheTemplatesForAGivenCriteriaV1(requestConfigurationTemplatesExportsTheTemplatesForAGivenCriteriaV1 *RequestConfigurationTemplatesExportsTheTemplatesForAGivenCriteriaV1) (*ResponseConfigurationTemplatesExportsTheTemplatesForAGivenCriteriaV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/template-programmer/template/exporttemplates"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestConfigurationTemplatesExportsTheTemplatesForAGivenCriteriaV1).
		SetResult(&ResponseConfigurationTemplatesExportsTheTemplatesForAGivenCriteriaV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.ExportsTheTemplatesForAGivenCriteriaV1(requestConfigurationTemplatesExportsTheTemplatesForAGivenCriteriaV1)
		}

		return nil, response, fmt.Errorf("error with operation ExportsTheTemplatesForAGivenCriteriaV1")
	}

	result := response.Result().(*ResponseConfigurationTemplatesExportsTheTemplatesForAGivenCriteriaV1)
	return result, response, err

}

//VersionTemplateV1 Version Template - f2a4-4a7e-4d5b-ab78
/* API to version the current contents of the template.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!version-template-v1
*/
func (s *ConfigurationTemplatesService) VersionTemplateV1(requestConfigurationTemplatesVersionTemplateV1 *RequestConfigurationTemplatesVersionTemplateV1) (*ResponseConfigurationTemplatesVersionTemplateV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/template-programmer/template/version"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestConfigurationTemplatesVersionTemplateV1).
		SetResult(&ResponseConfigurationTemplatesVersionTemplateV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.VersionTemplateV1(requestConfigurationTemplatesVersionTemplateV1)
		}

		return nil, response, fmt.Errorf("error with operation VersionTemplateV1")
	}

	result := response.Result().(*ResponseConfigurationTemplatesVersionTemplateV1)
	return result, response, err

}

//DeployTemplateV2 Deploy Template V2 - 02af-1bdf-4b48-9cbb
/* V2 API to deploy a template.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!deploy-template-v2
*/
func (s *ConfigurationTemplatesService) DeployTemplateV2(requestConfigurationTemplatesDeployTemplateV2 *RequestConfigurationTemplatesDeployTemplateV2) (*ResponseConfigurationTemplatesDeployTemplateV2, *resty.Response, error) {
	path := "/dna/intent/api/v2/template-programmer/template/deploy"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestConfigurationTemplatesDeployTemplateV2).
		SetResult(&ResponseConfigurationTemplatesDeployTemplateV2{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeployTemplateV2(requestConfigurationTemplatesDeployTemplateV2)
		}

		return nil, response, fmt.Errorf("error with operation DeployTemplateV2")
	}

	result := response.Result().(*ResponseConfigurationTemplatesDeployTemplateV2)
	return result, response, err

}

//UpdateProjectV1 Update Project - ecb8-8b89-4318-ac8d
/* This API is used to update an existing project.


 */
func (s *ConfigurationTemplatesService) UpdateProjectV1(requestConfigurationTemplatesUpdateProjectV1 *RequestConfigurationTemplatesUpdateProjectV1) (*ResponseConfigurationTemplatesUpdateProjectV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/template-programmer/project"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestConfigurationTemplatesUpdateProjectV1).
		SetResult(&ResponseConfigurationTemplatesUpdateProjectV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateProjectV1(requestConfigurationTemplatesUpdateProjectV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdateProjectV1")
	}

	result := response.Result().(*ResponseConfigurationTemplatesUpdateProjectV1)
	return result, response, err

}

//UpdateTemplateV1 Update Template - 2a80-39f5-4aab-86be
/* API to update a template.


 */
func (s *ConfigurationTemplatesService) UpdateTemplateV1(requestConfigurationTemplatesUpdateTemplateV1 *RequestConfigurationTemplatesUpdateTemplateV1) (*ResponseConfigurationTemplatesUpdateTemplateV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/template-programmer/template"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestConfigurationTemplatesUpdateTemplateV1).
		SetResult(&ResponseConfigurationTemplatesUpdateTemplateV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateTemplateV1(requestConfigurationTemplatesUpdateTemplateV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdateTemplateV1")
	}

	result := response.Result().(*ResponseConfigurationTemplatesUpdateTemplateV1)
	return result, response, err

}

//PreviewTemplateV1 Preview Template - 41bc-aaa6-4669-853e
/* API to preview a template.


 */
func (s *ConfigurationTemplatesService) PreviewTemplateV1(requestConfigurationTemplatesPreviewTemplateV1 *RequestConfigurationTemplatesPreviewTemplateV1) (*ResponseConfigurationTemplatesPreviewTemplateV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/template-programmer/template/preview"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestConfigurationTemplatesPreviewTemplateV1).
		SetResult(&ResponseConfigurationTemplatesPreviewTemplateV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.PreviewTemplateV1(requestConfigurationTemplatesPreviewTemplateV1)
		}
		return nil, response, fmt.Errorf("error with operation PreviewTemplateV1")
	}

	result := response.Result().(*ResponseConfigurationTemplatesPreviewTemplateV1)
	return result, response, err

}

//DeletesTheProjectV1 Deletes the project - 8cbb-79f4-4259-82d4
/* Deletes the project by its id


@param projectID projectId path parameter. projectId(UUID) of project to be deleted


Documentation Link: https://developer.cisco.com/docs/dna-center/#!deletes-the-project-v1
*/
func (s *ConfigurationTemplatesService) DeletesTheProjectV1(projectID string) (*ResponseConfigurationTemplatesDeletesTheProjectV1, *resty.Response, error) {
	//projectID string
	path := "/dna/intent/api/v1/template-programmer/project/{projectId}"
	path = strings.Replace(path, "{projectId}", fmt.Sprintf("%v", projectID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseConfigurationTemplatesDeletesTheProjectV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeletesTheProjectV1(projectID)
		}
		return nil, response, fmt.Errorf("error with operation DeletesTheProjectV1")
	}

	result := response.Result().(*ResponseConfigurationTemplatesDeletesTheProjectV1)
	return result, response, err

}

//DeletesTheTemplateV1 Deletes the template - a2bb-4afd-4699-8965
/* Deletes the template by its id


@param templateID templateId path parameter. templateId(UUID) of template to be deleted


Documentation Link: https://developer.cisco.com/docs/dna-center/#!deletes-the-template-v1
*/
func (s *ConfigurationTemplatesService) DeletesTheTemplateV1(templateID string) (*ResponseConfigurationTemplatesDeletesTheTemplateV1, *resty.Response, error) {
	//templateID string
	path := "/dna/intent/api/v1/template-programmer/template/{templateId}"
	path = strings.Replace(path, "{templateId}", fmt.Sprintf("%v", templateID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseConfigurationTemplatesDeletesTheTemplateV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeletesTheTemplateV1(templateID)
		}
		return nil, response, fmt.Errorf("error with operation DeletesTheTemplateV1")
	}

	result := response.Result().(*ResponseConfigurationTemplatesDeletesTheTemplateV1)
	return result, response, err

}

// Alias Function
func (s *ConfigurationTemplatesService) GetProjectsDetails(GetProjectsDetailsV2QueryParams *GetProjectsDetailsV2QueryParams) (*ResponseConfigurationTemplatesGetProjectsDetailsV2, *resty.Response, error) {
	return s.GetProjectsDetailsV2(GetProjectsDetailsV2QueryParams)
}

// Alias Function
func (s *ConfigurationTemplatesService) ImportsTheProjectsProvided(ImportsTheProjectsProvidedV1QueryParams *ImportsTheProjectsProvidedV1QueryParams) (*ResponseConfigurationTemplatesImportsTheProjectsProvidedV1, *resty.Response, error) {
	return s.ImportsTheProjectsProvidedV1(ImportsTheProjectsProvidedV1QueryParams)
}

// Alias Function
func (s *ConfigurationTemplatesService) CreateProject(requestConfigurationTemplatesCreateProjectV1 *RequestConfigurationTemplatesCreateProjectV1) (*ResponseConfigurationTemplatesCreateProjectV1, *resty.Response, error) {
	return s.CreateProjectV1(requestConfigurationTemplatesCreateProjectV1)
}

// Alias Function
func (s *ConfigurationTemplatesService) DeployTemplate(requestConfigurationTemplatesDeployTemplateV1 *RequestConfigurationTemplatesDeployTemplateV1) (*ResponseConfigurationTemplatesDeployTemplateV1, *resty.Response, error) {
	return s.DeployTemplateV1(requestConfigurationTemplatesDeployTemplateV1)
}

// Alias Function
func (s *ConfigurationTemplatesService) CreateTemplate(projectID string, requestConfigurationTemplatesCreateTemplateV1 *RequestConfigurationTemplatesCreateTemplateV1) (*ResponseConfigurationTemplatesCreateTemplateV1, *resty.Response, error) {
	return s.CreateTemplateV1(projectID, requestConfigurationTemplatesCreateTemplateV1)
}

// Alias Function
func (s *ConfigurationTemplatesService) UpdateProject(requestConfigurationTemplatesUpdateProjectV1 *RequestConfigurationTemplatesUpdateProjectV1) (*ResponseConfigurationTemplatesUpdateProjectV1, *resty.Response, error) {
	return s.UpdateProjectV1(requestConfigurationTemplatesUpdateProjectV1)
}

// Alias Function
func (s *ConfigurationTemplatesService) GetsTheDetailsOfAGivenProject(projectID string) (*ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectV1, *resty.Response, error) {
	return s.GetsTheDetailsOfAGivenProjectV1(projectID)
}

// Alias Function
func (s *ConfigurationTemplatesService) DeletesTheTemplate(templateID string) (*ResponseConfigurationTemplatesDeletesTheTemplateV1, *resty.Response, error) {
	return s.DeletesTheTemplateV1(templateID)
}

// Alias Function
func (s *ConfigurationTemplatesService) GetsTheTemplatesAvailable(GetsTheTemplatesAvailableV1QueryParams *GetsTheTemplatesAvailableV1QueryParams) (*ResponseConfigurationTemplatesGetsTheTemplatesAvailableV1, *resty.Response, error) {
	return s.GetsTheTemplatesAvailableV1(GetsTheTemplatesAvailableV1QueryParams)
}

// Alias Function
func (s *ConfigurationTemplatesService) ExportsTheProjectsForAGivenCriteria(requestConfigurationTemplatesExportsTheProjectsForAGivenCriteriaV1 *RequestConfigurationTemplatesExportsTheProjectsForAGivenCriteriaV1) (*ResponseConfigurationTemplatesExportsTheProjectsForAGivenCriteriaV1, *resty.Response, error) {
	return s.ExportsTheProjectsForAGivenCriteriaV1(requestConfigurationTemplatesExportsTheProjectsForAGivenCriteriaV1)
}

// Alias Function
func (s *ConfigurationTemplatesService) CreatesACloneOfTheGivenTemplate(name string, templateID string, projectID string, CreatesACloneOfTheGivenTemplateV1QueryParams *CreatesACloneOfTheGivenTemplateV1QueryParams) (*ResponseConfigurationTemplatesCreatesACloneOfTheGivenTemplateV1, *resty.Response, error) {
	return s.CreatesACloneOfTheGivenTemplateV1(name, templateID, projectID, CreatesACloneOfTheGivenTemplateV1QueryParams)
}

// Alias Function
func (s *ConfigurationTemplatesService) ImportsTheTemplatesProvided(projectName string, requestConfigurationTemplatesImportsTheTemplatesProvidedV1 *RequestConfigurationTemplatesImportsTheTemplatesProvidedV1, ImportsTheTemplatesProvidedV1QueryParams *ImportsTheTemplatesProvidedV1QueryParams) (*ResponseConfigurationTemplatesImportsTheTemplatesProvidedV1, *resty.Response, error) {
	return s.ImportsTheTemplatesProvidedV1(projectName, requestConfigurationTemplatesImportsTheTemplatesProvidedV1, ImportsTheTemplatesProvidedV1QueryParams)
}

// Alias Function
func (s *ConfigurationTemplatesService) PreviewTemplate(requestConfigurationTemplatesPreviewTemplateV1 *RequestConfigurationTemplatesPreviewTemplateV1) (*ResponseConfigurationTemplatesPreviewTemplateV1, *resty.Response, error) {
	return s.PreviewTemplateV1(requestConfigurationTemplatesPreviewTemplateV1)
}

// Alias Function
func (s *ConfigurationTemplatesService) DeletesTheProject(projectID string) (*ResponseConfigurationTemplatesDeletesTheProjectV1, *resty.Response, error) {
	return s.DeletesTheProjectV1(projectID)
}

// Alias Function
func (s *ConfigurationTemplatesService) ExportsTheTemplatesForAGivenCriteria(requestConfigurationTemplatesExportsTheTemplatesForAGivenCriteriaV1 *RequestConfigurationTemplatesExportsTheTemplatesForAGivenCriteriaV1) (*ResponseConfigurationTemplatesExportsTheTemplatesForAGivenCriteriaV1, *resty.Response, error) {
	return s.ExportsTheTemplatesForAGivenCriteriaV1(requestConfigurationTemplatesExportsTheTemplatesForAGivenCriteriaV1)
}

// Alias Function
func (s *ConfigurationTemplatesService) GetsAllTheVersionsOfAGivenTemplate(templateID string) (*ResponseConfigurationTemplatesGetsAllTheVersionsOfAGivenTemplateV1, *resty.Response, error) {
	return s.GetsAllTheVersionsOfAGivenTemplateV1(templateID)
}

// Alias Function
func (s *ConfigurationTemplatesService) StatusOfTemplateDeployment(deploymentID string) (*ResponseConfigurationTemplatesStatusOfTemplateDeploymentV1, *resty.Response, error) {
	return s.StatusOfTemplateDeploymentV1(deploymentID)
}

// Alias Function
func (s *ConfigurationTemplatesService) UpdateTemplate(requestConfigurationTemplatesUpdateTemplateV1 *RequestConfigurationTemplatesUpdateTemplateV1) (*ResponseConfigurationTemplatesUpdateTemplateV1, *resty.Response, error) {
	return s.UpdateTemplateV1(requestConfigurationTemplatesUpdateTemplateV1)
}

// Alias Function
func (s *ConfigurationTemplatesService) GetsAListOfProjects(GetsAListOfProjectsV1QueryParams *GetsAListOfProjectsV1QueryParams) (*ResponseConfigurationTemplatesGetsAListOfProjectsV1, *resty.Response, error) {
	return s.GetsAListOfProjectsV1(GetsAListOfProjectsV1QueryParams)
}

// Alias Function
func (s *ConfigurationTemplatesService) GetsDetailsOfAGivenTemplate(templateID string, GetsDetailsOfAGivenTemplateV1QueryParams *GetsDetailsOfAGivenTemplateV1QueryParams) (*ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateV1, *resty.Response, error) {
	return s.GetsDetailsOfAGivenTemplateV1(templateID, GetsDetailsOfAGivenTemplateV1QueryParams)
}

// Alias Function
func (s *ConfigurationTemplatesService) VersionTemplate(requestConfigurationTemplatesVersionTemplateV1 *RequestConfigurationTemplatesVersionTemplateV1) (*ResponseConfigurationTemplatesVersionTemplateV1, *resty.Response, error) {
	return s.VersionTemplateV1(requestConfigurationTemplatesVersionTemplateV1)
}
