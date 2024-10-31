package catalyst

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type SoftwareImageManagementSwimService service

type TriggerSoftwareImageActivationV1QueryParams struct {
	ScheduleValidate bool `url:"scheduleValidate,omitempty"` //scheduleValidate, validates data before schedule (Optional)
}
type TriggerSoftwareImageActivationV1HeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	ClientType  string `url:"Client-Type,omitempty"`  //Expects type string. Client-type (Optional)
	ClientURL   string `url:"Client-Url,omitempty"`   //Expects type string. Client-url (Optional)
}
type GetSoftwareImageDetailsV1QueryParams struct {
	ImageUUID            string `url:"imageUuid,omitempty"`            //imageUuid
	Name                 string `url:"name,omitempty"`                 //name
	Family               string `url:"family,omitempty"`               //family
	ApplicationType      string `url:"applicationType,omitempty"`      //applicationType
	ImageIntegrityStatus string `url:"imageIntegrityStatus,omitempty"` //imageIntegrityStatus - FAILURE, UNKNOWN, VERIFIED
	Version              string `url:"version,omitempty"`              //software Image Version
	ImageSeries          string `url:"imageSeries,omitempty"`          //image Series
	ImageName            string `url:"imageName,omitempty"`            //image Name
	IsTaggedGolden       bool   `url:"isTaggedGolden,omitempty"`       //is Tagged Golden
	IsCCORecommended     bool   `url:"isCCORecommended,omitempty"`     //is recommended from cisco.com
	IsCCOLatest          bool   `url:"isCCOLatest,omitempty"`          //is latest from cisco.com
	CreatedTime          int    `url:"createdTime,omitempty"`          //time in milliseconds (epoch format)
	ImageSizeGreaterThan int    `url:"imageSizeGreaterThan,omitempty"` //size in bytes
	ImageSizeLesserThan  int    `url:"imageSizeLesserThan,omitempty"`  //size in bytes
	SortBy               string `url:"sortBy,omitempty"`               //sort results by this field
	SortOrder            string `url:"sortOrder,omitempty"`            //sort order - 'asc' or 'des'. Default is asc
	Limit                int    `url:"limit,omitempty"`                //limit
	Offset               int    `url:"offset,omitempty"`               //offset
}
type ImportLocalSoftwareImageV1QueryParams struct {
	IsThirdParty              bool   `url:"isThirdParty,omitempty"`              //Third party Image check
	ThirdPartyVendor          string `url:"thirdPartyVendor,omitempty"`          //Third Party Vendor
	ThirdPartyImageFamily     string `url:"thirdPartyImageFamily,omitempty"`     //Third Party image family
	ThirdPartyApplicationType string `url:"thirdPartyApplicationType,omitempty"` //Third Party Application Type
}

type ImportLocalSoftwareImageMultipartFields struct {
	File     io.Reader
	FileName string
}

type ImportSoftwareImageViaURLV1QueryParams struct {
	ScheduleAt     string `url:"scheduleAt,omitempty"`     //Epoch Time (The number of milli-seconds since January 1 1970 UTC) at which the distribution should be scheduled (Optional)
	ScheduleDesc   string `url:"scheduleDesc,omitempty"`   //Custom Description (Optional)
	ScheduleOrigin string `url:"scheduleOrigin,omitempty"` //Originator of this call (Optional)
}
type ReturnsListOfSoftwareImagesV1QueryParams struct {
	SiteID                       string  `url:"siteId,omitempty"`                       //Site identifier to get the list of all available products under the site. The default value is the global site.  See https://developer.cisco.com/docs/dna-center/get-site for `siteId`
	ProductNameOrdinal           float64 `url:"productNameOrdinal,omitempty"`           //The product name ordinal is a unique value for each network device product. The productNameOrdinal can be obtained from the response of API `/dna/intent/api/v1/siteWiseProductNames`
	SupervisorProductNameOrdinal float64 `url:"supervisorProductNameOrdinal,omitempty"` //The supervisor engine module ordinal is a unique value for each supervisor module. The `supervisorProductNameOrdinal` can be obtained from the response of API `/dna/intent/api/v1/siteWiseProductNames`
	Imported                     bool    `url:"imported,omitempty"`                     //When the value is set to `true`, it will include physically imported images. Conversely, when the value is set to `false`, it will include image records from the cloud. The identifier for cloud images can be utilized to download images from Cisco.com to the disk.
	Name                         string  `url:"name,omitempty"`                         //Filter with software image or add-on name. Supports partial case-insensitive search. A minimum of 3 characters is required for the search.
	Version                      string  `url:"version,omitempty"`                      //Filter with image version. Supports partial case-insensitive search. A minimum of 3 characters is required for the search.
	Golden                       bool    `url:"golden,omitempty"`                       //When set to `true`, it will retrieve the images marked as tagged golden. When set to `false`, it will retrieve the images marked as not tagged golden.
	Integrity                    string  `url:"integrity,omitempty"`                    //Filter with verified images using Integrity Verification Available values: UNKNOWN, VERIFIED
	HasAddonImages               bool    `url:"hasAddonImages,omitempty"`               //When set to `true`, it will retrieve the images which have add-on images. When set to `false`, it will retrieve the images which do not have add-on images.
	IsAddonImages                bool    `url:"isAddonImages,omitempty"`                //When set to `true`, it will retrieve the images that an add-on image.  When set to `false`, it will retrieve the images that are not add-on images
	Offset                       float64 `url:"offset,omitempty"`                       //The first record to show for this page; the first record is numbered 1. The minimum value is 1.
	Limit                        float64 `url:"limit,omitempty"`                        //The number of records to show for this page. The minimum and maximum values are 1 and 500, respectively.
}
type ReturnsCountOfSoftwareImagesV1QueryParams struct {
	SiteID                       string  `url:"siteId,omitempty"`                       //Site identifier to get the list of all available products under the site. The default value is the global site.  See https://developer.cisco.com/docs/dna-center/get-site for siteId
	ProductNameOrdinal           float64 `url:"productNameOrdinal,omitempty"`           //The product name ordinal is a unique value for each network device product. The productNameOrdinal can be obtained from the response of the API `/dna/intent/api/v1/siteWiseProductNames`.
	SupervisorProductNameOrdinal float64 `url:"supervisorProductNameOrdinal,omitempty"` //The supervisor engine module ordinal is a unique value for each supervisor module. The `supervisorProductNameOrdinal` can be obtained from the response of API `/dna/intent/api/v1/siteWiseProductNames`
	Imported                     bool    `url:"imported,omitempty"`                     //When the value is set to `true`, it will include physically imported images. Conversely, when the value is set to `false`, it will include image records from the cloud. The identifier for cloud images can be utilised to download images from Cisco.com to the disk.
	Name                         string  `url:"name,omitempty"`                         //Filter with software image or add-on name. Supports partial case-insensitive search. A minimum of 3 characters is required for the search
	Version                      string  `url:"version,omitempty"`                      //Filter with image version. Supports partial case-insensitive search. A minimum of 3 characters is required for the search
	Golden                       string  `url:"golden,omitempty"`                       //When set to `true`, it will retrieve the images marked tagged golden. When set to `false`, it will retrieve the images marked not tagged golden.
	Integrity                    string  `url:"integrity,omitempty"`                    //Filter with verified images using Integrity Verification Available values: UNKNOWN, VERIFIED
	HasAddonImages               bool    `url:"hasAddonImages,omitempty"`               //When set to `true`, it will retrieve the images which have add-on images. When set to `false`, it will retrieve the images which do not have add-on images.
	IsAddonImages                bool    `url:"isAddonImages,omitempty"`                //When set to `true`, it will retrieve the images that an add-on image.  When set to `false`, it will retrieve the images that are not add-on images
}
type RetrievesNetworkDeviceProductNamesAssignedToASoftwareImageV1QueryParams struct {
	ProductName string  `url:"productName,omitempty"` //Filter with network device product name. Supports partial case-insensitive search. A minimum of 3 characters is required for the search.
	ProductID   string  `url:"productId,omitempty"`   //Filter with product ID (PID)
	Recommended string  `url:"recommended,omitempty"` //Filter with recommended source. If `CISCO` then the network device product assigned was recommended by Cisco and `USER` then the user has manually assigned. Available values: CISCO, USER
	Assigned    string  `url:"assigned,omitempty"`    //Filter with the assigned/unassigned, `ASSIGNED` option will filter network device products that are associated with the given image. The `NOT_ASSIGNED` option will filter network device products that have not yet been associated with the given image but apply to it. Available values: ASSIGNED, NOT_ASSIGNED
	Offset      float64 `url:"offset,omitempty"`      //The first record to show for this page; the first record is numbered 1. The minimum value is 1
	Limit       float64 `url:"limit,omitempty"`       //The number of records to show for this page. The minimum and maximum values are 1 and 500, respectively
}
type RetrievesTheCountOfAssignedNetworkDeviceProductsV1QueryParams struct {
	ProductName string `url:"productName,omitempty"` //Filter with network device product name. Supports partial case-insensitive search. A minimum of 3 characters are required for search.
	ProductID   string `url:"productId,omitempty"`   //Filter with product ID (PID)
	Recommended string `url:"recommended,omitempty"` //Filter with recommended source. If `CISCO` then the network device product assigned was recommended by Cisco and `USER` then the user has manually assigned. Available values : CISCO, USER
	Assigned    string `url:"assigned,omitempty"`    //Filter with the assigned/unassigned, `ASSIGNED` option will filter network device products that are associated with the given image. The `NOT_ASSIGNED` option will filter network device products that have not yet been associated with the given image but apply to it. Available values: ASSIGNED, NOT_ASSIGNED
}
type GetNetworkDeviceImageUpdatesV1QueryParams struct {
	ID                string  `url:"id,omitempty"`                //Update id which is unique for each network device under the parentId
	ParentID          string  `url:"parentId,omitempty"`          //Updates that have this parent id
	NetworkDeviceID   string  `url:"networkDeviceId,omitempty"`   //Network device id
	Status            string  `url:"status,omitempty"`            //Status of the image update. Available values : FAILURE, SUCCESS, IN_PROGRESS, PENDING
	ImageName         string  `url:"imageName,omitempty"`         //Software image name for the update
	HostName          string  `url:"hostName,omitempty"`          //Host name of the network device for the image update. Supports case-insensitive partial search
	ManagementAddress string  `url:"managementAddress,omitempty"` //Management address of the network device
	StartTime         float64 `url:"startTime,omitempty"`         //Image update started after the given time (as milliseconds since UNIX epoch)
	EndTime           float64 `url:"endTime,omitempty"`           //Image update started before the given time (as milliseconds since UNIX epoch)
	SortBy            string  `url:"sortBy,omitempty"`            //A property within the response to sort by.
	Order             string  `url:"order,omitempty"`             //Whether ascending or descending order should be used to sort the response.
	Offset            float64 `url:"offset,omitempty"`            //The first record to show for this page; the first record is numbered 1.
	Limit             float64 `url:"limit,omitempty"`             //The number of records to show for this page.
}
type CountOfNetworkDeviceImageUpdatesV1QueryParams struct {
	ID                string  `url:"id,omitempty"`                //Update id which is unique for each network device under the parentId
	ParentID          string  `url:"parentId,omitempty"`          //Updates that have this parent id
	NetworkDeviceID   string  `url:"networkDeviceId,omitempty"`   //Network device id
	Status            string  `url:"status,omitempty"`            //Status of the image update. Available values: FAILURE, SUCCESS, IN_PROGRESS, PENDING
	ImageName         string  `url:"imageName,omitempty"`         //Software image name for the update
	HostName          string  `url:"hostName,omitempty"`          //Host name of the network device for the image update. Supports case-insensitive partial search.
	ManagementAddress string  `url:"managementAddress,omitempty"` //Management address of the network device
	StartTime         float64 `url:"startTime,omitempty"`         //Image update started after the given time (as milliseconds since UNIX epoch).
	EndTime           float64 `url:"endTime,omitempty"`           //Image update started before the given time (as milliseconds since UNIX epoch).
}
type RetrievesTheListOfNetworkDeviceProductNamesV1QueryParams struct {
	ProductName string  `url:"productName,omitempty"` //Filter with network device product name. Supports partial case-insensitive search. A minimum of 3 characters are required for search
	ProductID   string  `url:"productId,omitempty"`   //Filter with product ID (PID)
	Offset      float64 `url:"offset,omitempty"`      //The first record to show for this page; the first record is numbered 1. The minimum value is 1.
	Limit       float64 `url:"limit,omitempty"`       //The number of records to show for this page. The minimum and maximum values are 1 and 500, respectively.
}
type CountOfNetworkProductNamesV1QueryParams struct {
	ProductName string `url:"productName,omitempty"` //Filter with network device product name. Supports partial case-insensitive search. A minimum of 3 characters are required for search
	ProductID   string `url:"productId,omitempty"`   //Filter with product ID (PID)
}
type ReturnsNetworkDeviceProductNamesForASiteV1QueryParams struct {
	SiteID      string  `url:"siteId,omitempty"`      //Site identifier to get the list of all available products under the site. The default value is the global site.  See https://developer.cisco.com/docs/dna-center/get-site for siteId
	ProductName string  `url:"productName,omitempty"` //Filter with network device product name. Supports partial case-insensitive search. A minimum of 3 characters are required for search
	Offset      float64 `url:"offset,omitempty"`      //The first record to show for this page; the first record is numbered 1. The minimum value is 1
	Limit       float64 `url:"limit,omitempty"`       //The number of records to show for this page. The minimum and maximum values are 1 and 500, respectively
}
type ReturnsTheCountOfNetworkDeviceProductNamesForASiteV1QueryParams struct {
	SiteID      string `url:"siteId,omitempty"`      //Site identifier to get the list of all available products under the site. The default value is global site id. See https://developer.cisco.com/docs/dna-center/get-site/ for siteId
	ProductName string `url:"productName,omitempty"` //Filter with network device product name. Supports partial case-insensitive search. A minimum of 3 characters are required for search
}

type ResponseSoftwareImageManagementSwimTriggerSoftwareImageActivationV1 struct {
	Response *ResponseSoftwareImageManagementSwimTriggerSoftwareImageActivationV1Response `json:"response,omitempty"` //
	Version  string                                                                       `json:"version,omitempty"`  //
}
type ResponseSoftwareImageManagementSwimTriggerSoftwareImageActivationV1Response struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseSoftwareImageManagementSwimTriggerSoftwareImageDistributionV1 struct {
	Response *ResponseSoftwareImageManagementSwimTriggerSoftwareImageDistributionV1Response `json:"response,omitempty"` //
	Version  string                                                                         `json:"version,omitempty"`  //
}
type ResponseSoftwareImageManagementSwimTriggerSoftwareImageDistributionV1Response struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseSoftwareImageManagementSwimGetSoftwareImageDetailsV1 struct {
	Response *[]ResponseSoftwareImageManagementSwimGetSoftwareImageDetailsV1Response `json:"response,omitempty"` //
	Version  string                                                                  `json:"version,omitempty"`  //
}
type ResponseSoftwareImageManagementSwimGetSoftwareImageDetailsV1Response struct {
	ApplicableDevicesForImage *[]ResponseSoftwareImageManagementSwimGetSoftwareImageDetailsV1ResponseApplicableDevicesForImage `json:"applicableDevicesForImage,omitempty"` //
	ApplicationType           string                                                                                           `json:"applicationType,omitempty"`           //
	CreatedTime               string                                                                                           `json:"createdTime,omitempty"`               //
	ExtendedAttributes        *ResponseSoftwareImageManagementSwimGetSoftwareImageDetailsV1ResponseExtendedAttributes          `json:"extendedAttributes,omitempty"`        //
	Family                    string                                                                                           `json:"family,omitempty"`                    //
	Feature                   string                                                                                           `json:"feature,omitempty"`                   //
	FileServiceID             string                                                                                           `json:"fileServiceId,omitempty"`             //
	FileSize                  string                                                                                           `json:"fileSize,omitempty"`                  //
	ImageIntegrityStatus      string                                                                                           `json:"imageIntegrityStatus,omitempty"`      //
	ImageName                 string                                                                                           `json:"imageName,omitempty"`                 //
	ImageSeries               []string                                                                                         `json:"imageSeries,omitempty"`               //
	ImageSource               string                                                                                           `json:"imageSource,omitempty"`               //
	ImageType                 string                                                                                           `json:"imageType,omitempty"`                 //
	ImageUUID                 string                                                                                           `json:"imageUuid,omitempty"`                 //
	ImportSourceType          string                                                                                           `json:"importSourceType,omitempty"`          //
	IsTaggedGolden            *bool                                                                                            `json:"isTaggedGolden,omitempty"`            //
	Md5Checksum               string                                                                                           `json:"md5Checksum,omitempty"`               //
	Name                      string                                                                                           `json:"name,omitempty"`                      //
	ProfileInfo               *[]ResponseSoftwareImageManagementSwimGetSoftwareImageDetailsV1ResponseProfileInfo               `json:"profileInfo,omitempty"`               //
	ShaCheckSum               string                                                                                           `json:"shaCheckSum,omitempty"`               //
	Vendor                    string                                                                                           `json:"vendor,omitempty"`                    //
	Version                   string                                                                                           `json:"version,omitempty"`                   //
}
type ResponseSoftwareImageManagementSwimGetSoftwareImageDetailsV1ResponseApplicableDevicesForImage struct {
	MdfID       string   `json:"mdfId,omitempty"`       //
	ProductID   []string `json:"productId,omitempty"`   //
	ProductName string   `json:"productName,omitempty"` //
}
type ResponseSoftwareImageManagementSwimGetSoftwareImageDetailsV1ResponseExtendedAttributes interface{}
type ResponseSoftwareImageManagementSwimGetSoftwareImageDetailsV1ResponseProfileInfo struct {
	Description        string                                                                                             `json:"description,omitempty"`        //
	ExtendedAttributes *ResponseSoftwareImageManagementSwimGetSoftwareImageDetailsV1ResponseProfileInfoExtendedAttributes `json:"extendedAttributes,omitempty"` //
	Memory             *int                                                                                               `json:"memory,omitempty"`             //
	ProductType        string                                                                                             `json:"productType,omitempty"`        //
	ProfileName        string                                                                                             `json:"profileName,omitempty"`        //
	Shares             *int                                                                                               `json:"shares,omitempty"`             //
	VCPU               *int                                                                                               `json:"vCpu,omitempty"`               //
}
type ResponseSoftwareImageManagementSwimGetSoftwareImageDetailsV1ResponseProfileInfoExtendedAttributes interface{}
type ResponseSoftwareImageManagementSwimGetDeviceFamilyIDentifiersV1 struct {
	Version  string                                                                     `json:"version,omitempty"`  // Response Version
	Response *[]ResponseSoftwareImageManagementSwimGetDeviceFamilyIDentifiersV1Response `json:"response,omitempty"` //
}
type ResponseSoftwareImageManagementSwimGetDeviceFamilyIDentifiersV1Response struct {
	DeviceFamily           string `json:"deviceFamily,omitempty"`           // Device Family e.g. : Cisco Catalyst 6503 Switch-Catalyst 6500 Series Supervisor Engine 2T
	DeviceFamilyIDentifier string `json:"deviceFamilyIdentifier,omitempty"` // Device Family Identifier used for tagging an image golden for certain Device Family e.g. : 277696480-283933147
}
type ResponseSoftwareImageManagementSwimTagAsGoldenImageV1 struct {
	Version  string                                                         `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseSoftwareImageManagementSwimTagAsGoldenImageV1Response `json:"response,omitempty"` //
}
type ResponseSoftwareImageManagementSwimTagAsGoldenImageV1Response struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseSoftwareImageManagementSwimRemoveGoldenTagForImageV1 struct {
	Version  string                                                                `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseSoftwareImageManagementSwimRemoveGoldenTagForImageV1Response `json:"response,omitempty"` //
}
type ResponseSoftwareImageManagementSwimRemoveGoldenTagForImageV1Response struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseSoftwareImageManagementSwimGetGoldenTagStatusOfAnImageV1 struct {
	Version  string                                                                    `json:"version,omitempty"`  // Response Version. E.G. : 1.0
	Response *ResponseSoftwareImageManagementSwimGetGoldenTagStatusOfAnImageV1Response `json:"response,omitempty"` //
}
type ResponseSoftwareImageManagementSwimGetGoldenTagStatusOfAnImageV1Response struct {
	DeviceRole        string `json:"deviceRole,omitempty"`        // Device Role. Possible Values : ALL, UNKNOWN, ACCESS, BORDER ROUTER, DISTRIBUTION and CORE.
	TaggedGolden      *bool  `json:"taggedGolden,omitempty"`      // Tagged Golden.
	InheritedSiteName string `json:"inheritedSiteName,omitempty"` // Inherited Site Name. If the Golden Tag is not tagged for the current site but is inherited from a higher enclosing site, it will contain the name of the site from where the tag is inherited.
	InheritedSiteID   string `json:"inheritedSiteId,omitempty"`   // Inherited Site Id. If the Golden Tag is not tagged for the current site but is inherited from a higher enclosing site, it will contain the uuid of the site from where the tag is inherited. In case the golden tag is inherited from the Global site the value will be "-1".
}
type ResponseSoftwareImageManagementSwimImportLocalSoftwareImageV1 struct {
	Response *ResponseSoftwareImageManagementSwimImportLocalSoftwareImageV1Response `json:"response,omitempty"` //
	Version  string                                                                 `json:"version,omitempty"`  //
}
type ResponseSoftwareImageManagementSwimImportLocalSoftwareImageV1Response struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseSoftwareImageManagementSwimImportSoftwareImageViaURLV1 struct {
	Response *ResponseSoftwareImageManagementSwimImportSoftwareImageViaURLV1Response `json:"response,omitempty"` //
	Version  string                                                                  `json:"version,omitempty"`  //
}
type ResponseSoftwareImageManagementSwimImportSoftwareImageViaURLV1Response struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseSoftwareImageManagementSwimReturnsListOfSoftwareImagesV1 struct {
	Response *[]ResponseSoftwareImageManagementSwimReturnsListOfSoftwareImagesV1Response `json:"response,omitempty"` //
	Version  string                                                                      `json:"version,omitempty"`  // Version
}
type ResponseSoftwareImageManagementSwimReturnsListOfSoftwareImagesV1Response struct {
	ID                   string                                                                                          `json:"id,omitempty"`                   // Software image identifier
	Imported             *bool                                                                                           `json:"imported,omitempty"`             // Flag for image info whether it is imported image or cloud image
	Name                 string                                                                                          `json:"name,omitempty"`                 // Name of the software image
	Version              string                                                                                          `json:"version,omitempty"`              // Software image  version
	ImageType            string                                                                                          `json:"imageType,omitempty"`            // Software image type
	Recommended          string                                                                                          `json:"recommended,omitempty"`          // CISCO if the image is recommended from Cisco.com
	CiscoLatest          *bool                                                                                           `json:"ciscoLatest,omitempty"`          // `true` if the image is latest/suggested from Cisco.com
	IntegrityStatus      string                                                                                          `json:"integrityStatus,omitempty"`      // Image Integrity verification status with Known Good Verification
	IsAddonImage         *bool                                                                                           `json:"isAddonImage,omitempty"`         // The value of `true` will indicate the image as an add-on image, while the value of `false` will indicate software image
	HasAddonImages       *bool                                                                                           `json:"hasAddonImages,omitempty"`       // Software images that have an applicable list of add-on images. The value of `true` will return software images with add-on images, while the value of `false` will return software images without add-on images
	GoldenTaggingDetails *[]ResponseSoftwareImageManagementSwimReturnsListOfSoftwareImagesV1ResponseGoldenTaggingDetails `json:"goldenTaggingDetails,omitempty"` //
	ProductNames         *[]ResponseSoftwareImageManagementSwimReturnsListOfSoftwareImagesV1ResponseProductNames         `json:"productNames,omitempty"`         //
	IsGoldenTagged       *bool                                                                                           `json:"isGoldenTagged,omitempty"`       // The value of `true` will indicate the image marked as golden, while the value of `false` will indicate the image not marked as golden
}
type ResponseSoftwareImageManagementSwimReturnsListOfSoftwareImagesV1ResponseGoldenTaggingDetails struct {
	DeviceRoles       []string `json:"deviceRoles,omitempty"`       // Golden tagging based on the device roles
	DeviceTags        []string `json:"deviceTags,omitempty"`        // Golden tagging based on the device tags
	InheritedSiteID   string   `json:"inheritedSiteId,omitempty"`   // The Site Id of the site that this setting is inherited from.
	InheritedSiteName string   `json:"inheritedSiteName,omitempty"` // The name of the site that this setting is inherited from
}
type ResponseSoftwareImageManagementSwimReturnsListOfSoftwareImagesV1ResponseProductNames struct {
	ID                           string   `json:"id,omitempty"`                           // Product name ordinal is unique value for each network device product
	ProductName                  string   `json:"productName,omitempty"`                  // Network device product name
	ProductNameOrdinal           *float64 `json:"productNameOrdinal,omitempty"`           // Product name ordinal is unique value for each network device product
	SupervisorProductName        string   `json:"supervisorProductName,omitempty"`        // Name of the Supervisor Engine Module, supported by the `productName`.                  Example: The `Cisco Catalyst 9404R Switch` chassis is capable of supporting  different supervisor engine modules: the `Cisco Catalyst 9400 Supervisor Engine-1`, the `Cisco Catalyst 9400 Supervisor Engine-1XL`, the `Cisco Catalyst 9400 Supervisor Engine-1XL-Y`, etc.
	SupervisorProductNameOrdinal *float64 `json:"supervisorProductNameOrdinal,omitempty"` // Supervisor Engine Module Ordinal, supported by the `productNameOrdinal`. Example: The `286315691` chassis ordinal is capable of supporting different supervisor engine module ordinals: `286316172`, `286316710`, `286320394` etc.
}
type ResponseSoftwareImageManagementSwimReturnsCountOfSoftwareImagesV1 struct {
	Response *ResponseSoftwareImageManagementSwimReturnsCountOfSoftwareImagesV1Response `json:"response,omitempty"` //
	Version  string                                                                     `json:"version,omitempty"`  // The version of the response
}
type ResponseSoftwareImageManagementSwimReturnsCountOfSoftwareImagesV1Response struct {
	Count *int `json:"count,omitempty"` // Reports a count, for example, a total count of records in a given resource.
}
type ResponseSoftwareImageManagementSwimAddImageDistributionServerV1 struct {
	Response *ResponseSoftwareImageManagementSwimAddImageDistributionServerV1Response `json:"response,omitempty"` //
	Version  string                                                                   `json:"version,omitempty"`  // Response version
}
type ResponseSoftwareImageManagementSwimAddImageDistributionServerV1Response struct {
	TaskID string `json:"taskId,omitempty"` // The UUID of the task
	URL    string `json:"url,omitempty"`    // The path to the API endpoint to GET for information on the task
}
type ResponseSoftwareImageManagementSwimRetrieveImageDistributionServersV1 struct {
	Response *[]ResponseSoftwareImageManagementSwimRetrieveImageDistributionServersV1Response `json:"response,omitempty"` //
	Version  string                                                                           `json:"version,omitempty"`  // Response version
}
type ResponseSoftwareImageManagementSwimRetrieveImageDistributionServersV1Response struct {
	ID            string   `json:"id,omitempty"`            // Unique identifier for the server
	Username      string   `json:"username,omitempty"`      // Server username
	ServerAddress string   `json:"serverAddress,omitempty"` // FQDN or IP address of the server
	PortNumber    *float64 `json:"portNumber,omitempty"`    // Port number
	RootLocation  string   `json:"rootLocation,omitempty"`  // Server root location
}
type ResponseSoftwareImageManagementSwimUpdateRemoteImageDistributionServerV1 struct {
	Response *ResponseSoftwareImageManagementSwimUpdateRemoteImageDistributionServerV1Response `json:"response,omitempty"` //
	Version  string                                                                            `json:"version,omitempty"`  // Response version
}
type ResponseSoftwareImageManagementSwimUpdateRemoteImageDistributionServerV1Response struct {
	TaskID string `json:"taskId,omitempty"` // The UUID of the task
	URL    string `json:"url,omitempty"`    // The path to the API endpoint to GET for information on the task
}
type ResponseSoftwareImageManagementSwimRetrieveSpecificImageDistributionServerV1 struct {
	Response *ResponseSoftwareImageManagementSwimRetrieveSpecificImageDistributionServerV1Response `json:"response,omitempty"` //
	Version  string                                                                                `json:"version,omitempty"`  // Response version
}
type ResponseSoftwareImageManagementSwimRetrieveSpecificImageDistributionServerV1Response struct {
	ID            string   `json:"id,omitempty"`            // Unique identifier for the server
	ServerAddress string   `json:"serverAddress,omitempty"` // FQDN or IP address of the server
	PortNumber    *float64 `json:"portNumber,omitempty"`    // Port number
	RootLocation  string   `json:"rootLocation,omitempty"`  // Server root location
	Username      string   `json:"username,omitempty"`      // Server username
}
type ResponseSoftwareImageManagementSwimRemoveImageDistributionServerV1 struct {
	Response *ResponseSoftwareImageManagementSwimRemoveImageDistributionServerV1Response `json:"response,omitempty"` //
	Version  string                                                                      `json:"version,omitempty"`  // Response version
}
type ResponseSoftwareImageManagementSwimRemoveImageDistributionServerV1Response struct {
	TaskID string `json:"taskId,omitempty"` // The UUID of the task
	URL    string `json:"url,omitempty"`    // The path to the API endpoint to GET for information on the task
}
type ResponseSoftwareImageManagementSwimRetrieveApplicableAddOnImagesForTheGivenSoftwareImageV1 struct {
	Response *[]ResponseSoftwareImageManagementSwimRetrieveApplicableAddOnImagesForTheGivenSoftwareImageV1Response `json:"response,omitempty"` //
	Version  string                                                                                                `json:"version,omitempty"`  // Version
}
type ResponseSoftwareImageManagementSwimRetrieveApplicableAddOnImagesForTheGivenSoftwareImageV1Response struct {
	ID                   string                                                                                                                    `json:"id,omitempty"`                   // Software image identifier
	Imported             *bool                                                                                                                     `json:"imported,omitempty"`             // Flag for image info whether it is imported image or cloud image
	Name                 string                                                                                                                    `json:"name,omitempty"`                 // Name of the software image
	Version              string                                                                                                                    `json:"version,omitempty"`              // Software image  version
	ImageType            string                                                                                                                    `json:"imageType,omitempty"`            // Software image type
	Recommended          string                                                                                                                    `json:"recommended,omitempty"`          // CISCO if the image is recommended from Cisco.com
	CiscoLatest          *bool                                                                                                                     `json:"ciscoLatest,omitempty"`          // `true` if the image is latest/suggested from Cisco.com
	IntegrityStatus      string                                                                                                                    `json:"integrityStatus,omitempty"`      // Image Integrity verification status with Known Good Verification
	IsAddonImage         *bool                                                                                                                     `json:"isAddonImage,omitempty"`         // The value of `true` will indicate the image as an add-on image, while the value of `false` will indicate software image
	HasAddonImages       *bool                                                                                                                     `json:"hasAddonImages,omitempty"`       // Software images that have an applicable list of add-on images. The value of `true` will return software images with add-on images, while the value of `false` will return software images without add-on images
	GoldenTaggingDetails *[]ResponseSoftwareImageManagementSwimRetrieveApplicableAddOnImagesForTheGivenSoftwareImageV1ResponseGoldenTaggingDetails `json:"goldenTaggingDetails,omitempty"` //
	ProductNames         *[]ResponseSoftwareImageManagementSwimRetrieveApplicableAddOnImagesForTheGivenSoftwareImageV1ResponseProductNames         `json:"productNames,omitempty"`         //
	IsGoldenTagged       *bool                                                                                                                     `json:"isGoldenTagged,omitempty"`       // The value of `true` will indicate the image marked as golden, while the value of `false` will indicate the image not marked as golden
}
type ResponseSoftwareImageManagementSwimRetrieveApplicableAddOnImagesForTheGivenSoftwareImageV1ResponseGoldenTaggingDetails struct {
	DeviceRoles       []string `json:"deviceRoles,omitempty"`       // Golden tagging based on the device roles
	DeviceTags        []string `json:"deviceTags,omitempty"`        // Golden tagging based on the device tags
	InheritedSiteID   string   `json:"inheritedSiteId,omitempty"`   // The Site Id of the site that this setting is inherited from.
	InheritedSiteName string   `json:"inheritedSiteName,omitempty"` // The name of the site that this setting is inherited from
}
type ResponseSoftwareImageManagementSwimRetrieveApplicableAddOnImagesForTheGivenSoftwareImageV1ResponseProductNames struct {
	ID                           string   `json:"id,omitempty"`                           // Product name ordinal is unique value for each network device product
	ProductName                  string   `json:"productName,omitempty"`                  // Network device product name
	ProductNameOrdinal           *float64 `json:"productNameOrdinal,omitempty"`           // Product name ordinal is unique value for each network device product
	SupervisorProductName        string   `json:"supervisorProductName,omitempty"`        // Name of the Supervisor Engine Module, supported by the `productName`.                  Example: The `Cisco Catalyst 9404R Switch` chassis is capable of supporting  different supervisor engine modules: the `Cisco Catalyst 9400 Supervisor Engine-1`, the `Cisco Catalyst 9400 Supervisor Engine-1XL`, the `Cisco Catalyst 9400 Supervisor Engine-1XL-Y`, etc.
	SupervisorProductNameOrdinal *float64 `json:"supervisorProductNameOrdinal,omitempty"` // Supervisor Engine Module Ordinal, supported by the `productNameOrdinal`. Example: The `286315691` chassis ordinal is capable of supporting different supervisor engine module ordinals: `286316172`, `286316710`, `286320394` etc.
}
type ResponseSoftwareImageManagementSwimReturnsCountOfAddOnImagesV1 struct {
	Response *ResponseSoftwareImageManagementSwimReturnsCountOfAddOnImagesV1Response `json:"response,omitempty"` //
	Version  string                                                                  `json:"version,omitempty"`  // The version of the response
}
type ResponseSoftwareImageManagementSwimReturnsCountOfAddOnImagesV1Response struct {
	Count *int `json:"count,omitempty"` // Reports a count, for example, a total count of records in a given resource.
}
type ResponseSoftwareImageManagementSwimDownloadTheSoftwareImageV1 struct {
	Response *ResponseSoftwareImageManagementSwimDownloadTheSoftwareImageV1Response `json:"response,omitempty"` //
	Version  string                                                                 `json:"version,omitempty"`  // Response version
}
type ResponseSoftwareImageManagementSwimDownloadTheSoftwareImageV1Response struct {
	TaskID string `json:"taskId,omitempty"` // The UUID of the task
	URL    string `json:"url,omitempty"`    // The path to the API endpoint to GET for information on the task
}
type ResponseSoftwareImageManagementSwimAssignNetworkDeviceProductNameToTheGivenSoftwareImageV1 struct {
	Response *ResponseSoftwareImageManagementSwimAssignNetworkDeviceProductNameToTheGivenSoftwareImageV1Response `json:"response,omitempty"` //
	Version  string                                                                                              `json:"version,omitempty"`  // Response version
}
type ResponseSoftwareImageManagementSwimAssignNetworkDeviceProductNameToTheGivenSoftwareImageV1Response struct {
	TaskID string `json:"taskId,omitempty"` // The UUID of the task
	URL    string `json:"url,omitempty"`    // The path to the API endpoint to GET for information on the task
}
type ResponseSoftwareImageManagementSwimRetrievesNetworkDeviceProductNamesAssignedToASoftwareImageV1 struct {
	Response *[]ResponseSoftwareImageManagementSwimRetrievesNetworkDeviceProductNamesAssignedToASoftwareImageV1Response `json:"response,omitempty"` //
	Version  string                                                                                                     `json:"version,omitempty"`  // Response version
}
type ResponseSoftwareImageManagementSwimRetrievesNetworkDeviceProductNamesAssignedToASoftwareImageV1Response struct {
	ID                 string   `json:"id,omitempty"`                 // Product name ordinal is unique value for each network device product
	ProductName        string   `json:"productName,omitempty"`        // Network device product name
	ProductNameOrdinal *float64 `json:"productNameOrdinal,omitempty"` // Product name ordinal is unique value for each network device product
	ProductIDs         []string `json:"productIds,omitempty"`         // Supported PIDs
	SiteIDs            []string `json:"siteIds,omitempty"`            // Sites where all  this image is assigned
	Recommended        string   `json:"recommended,omitempty"`        // If 'CISCO' network device product recommandation came from Cisco.com and 'USER' manually assigned
}
type ResponseSoftwareImageManagementSwimRetrievesTheCountOfAssignedNetworkDeviceProductsV1 struct {
	Response *ResponseSoftwareImageManagementSwimRetrievesTheCountOfAssignedNetworkDeviceProductsV1Response `json:"response,omitempty"` //
	Version  string                                                                                         `json:"version,omitempty"`  // The version of the response
}
type ResponseSoftwareImageManagementSwimRetrievesTheCountOfAssignedNetworkDeviceProductsV1Response struct {
	Count *int `json:"count,omitempty"` // Reports a count, for example, a total count of records in a given resource.
}
type ResponseSoftwareImageManagementSwimUnassignNetworkDeviceProductNameFromTheGivenSoftwareImageV1 struct {
	Response *ResponseSoftwareImageManagementSwimUnassignNetworkDeviceProductNameFromTheGivenSoftwareImageV1Response `json:"response,omitempty"` //
	Version  string                                                                                                  `json:"version,omitempty"`  // Response version
}
type ResponseSoftwareImageManagementSwimUnassignNetworkDeviceProductNameFromTheGivenSoftwareImageV1Response struct {
	TaskID string `json:"taskId,omitempty"` // The UUID of the task
	URL    string `json:"url,omitempty"`    // The path to the API endpoint to GET for information on the task
}
type ResponseSoftwareImageManagementSwimUpdateTheListOfSitesForTheNetworkDeviceProductNameAssignedToTheSoftwareImageV1 struct {
	Response *ResponseSoftwareImageManagementSwimUpdateTheListOfSitesForTheNetworkDeviceProductNameAssignedToTheSoftwareImageV1Response `json:"response,omitempty"` //
	Version  string                                                                                                                     `json:"version,omitempty"`  // Response version
}
type ResponseSoftwareImageManagementSwimUpdateTheListOfSitesForTheNetworkDeviceProductNameAssignedToTheSoftwareImageV1Response struct {
	TaskID string `json:"taskId,omitempty"` // The UUID of the task
	URL    string `json:"url,omitempty"`    // The path to the API endpoint to GET for information on the task
}
type ResponseSoftwareImageManagementSwimGetNetworkDeviceImageUpdatesV1 struct {
	Response *[]ResponseSoftwareImageManagementSwimGetNetworkDeviceImageUpdatesV1Response `json:"response,omitempty"` //
	Version  string                                                                       `json:"version,omitempty"`  // API response version
}
type ResponseSoftwareImageManagementSwimGetNetworkDeviceImageUpdatesV1Response struct {
	ID                 string   `json:"id,omitempty"`                 // Unique identifier for the image update
	ParentID           string   `json:"parentId,omitempty"`           // Parent identifier for the image update
	StartTime          *float64 `json:"startTime,omitempty"`          // Image update started after the given time (as milliseconds since UNIX epoch)
	EndTime            *float64 `json:"endTime,omitempty"`            // Image update end time (as milliseconds since UNIX epoch)
	Status             string   `json:"status,omitempty"`             // Status of the image update
	NetworkDeviceID    string   `json:"networkDeviceId,omitempty"`    // Network device identifier
	ManagementAddress  string   `json:"managementAddress,omitempty"`  // Management address of the network device
	HostName           string   `json:"hostName,omitempty"`           // Host name of the network device for the image update
	UpdateImageVersion string   `json:"updateImageVersion,omitempty"` // Software image version
	Type               string   `json:"type,omitempty"`               // Type of the image update
}
type ResponseSoftwareImageManagementSwimCountOfNetworkDeviceImageUpdatesV1 struct {
	Response *ResponseSoftwareImageManagementSwimCountOfNetworkDeviceImageUpdatesV1Response `json:"response,omitempty"` //
	Version  string                                                                         `json:"version,omitempty"`  // The version of the response
}
type ResponseSoftwareImageManagementSwimCountOfNetworkDeviceImageUpdatesV1Response struct {
	Count *int `json:"count,omitempty"` // Reports a count, for example, a total count of records in a given resource.
}
type ResponseSoftwareImageManagementSwimRetrievesTheListOfNetworkDeviceProductNamesV1 struct {
	Response *[]ResponseSoftwareImageManagementSwimRetrievesTheListOfNetworkDeviceProductNamesV1Response `json:"response,omitempty"` //
	Version  string                                                                                      `json:"version,omitempty"`  // Response version
}
type ResponseSoftwareImageManagementSwimRetrievesTheListOfNetworkDeviceProductNamesV1Response struct {
	ID                 string   `json:"id,omitempty"`                 // Product name ordinal is unique value for each network device product
	ProductName        string   `json:"productName,omitempty"`        // Network device product name
	ProductNameOrdinal *float64 `json:"productNameOrdinal,omitempty"` // Product name ordinal is unique value for each network device product
	ProductIDs         []string `json:"productIds,omitempty"`         // Supported PIDs
}
type ResponseSoftwareImageManagementSwimCountOfNetworkProductNamesV1 struct {
	Response *ResponseSoftwareImageManagementSwimCountOfNetworkProductNamesV1Response `json:"response,omitempty"` //
	Version  string                                                                   `json:"version,omitempty"`  // The version of the response
}
type ResponseSoftwareImageManagementSwimCountOfNetworkProductNamesV1Response struct {
	Count *int `json:"count,omitempty"` // Reports a count, for example, a total count of records in a given resource.
}
type ResponseSoftwareImageManagementSwimRetrieveNetworkDeviceProductNameV1 struct {
	Response *ResponseSoftwareImageManagementSwimRetrieveNetworkDeviceProductNameV1Response `json:"response,omitempty"` //
	Version  string                                                                         `json:"version,omitempty"`  // Response version
}
type ResponseSoftwareImageManagementSwimRetrieveNetworkDeviceProductNameV1Response struct {
	ID                 string   `json:"id,omitempty"`                 // Product name ordinal is unique value for each network device product
	ProductName        string   `json:"productName,omitempty"`        // Network device product name
	ProductNameOrdinal *float64 `json:"productNameOrdinal,omitempty"` // Product name ordinal is unique value for each network device product
	ProductIDs         []string `json:"productIds,omitempty"`         // Supported PIDs
}
type ResponseSoftwareImageManagementSwimReturnsNetworkDeviceProductNamesForASiteV1 struct {
	Response *ResponseSoftwareImageManagementSwimReturnsNetworkDeviceProductNamesForASiteV1Response `json:"response,omitempty"` //
	Version  string                                                                                 `json:"version,omitempty"`  // Response version
}
type ResponseSoftwareImageManagementSwimReturnsNetworkDeviceProductNamesForASiteV1Response struct {
	ID                           string                                                                                             `json:"id,omitempty"`                           // The unique identifier for the record is the `id`. If there is no supervisor engine involved, the `id` will be the same as the `productNameOrdinal`. However, if the supervisor engine is applicable, the `id` will be in the form of `<productNameOrdinal>-<supervisorProductNameOrdinal>`.
	ProductNameOrdinal           *float64                                                                                           `json:"productNameOrdinal,omitempty"`           // Product name ordinal
	ProductName                  string                                                                                             `json:"productName,omitempty"`                  // Name of product
	SupervisorProductName        string                                                                                             `json:"supervisorProductName,omitempty"`        // Name of the Supervisor Engine Module, supported by the `productName`.                  Example: The `Cisco Catalyst 9404R Switch` chassis is capable of supporting  different supervisor engine modules: the `Cisco Catalyst 9400 Supervisor Engine-1`, the `Cisco Catalyst 9400 Supervisor Engine-1XL`, the `Cisco Catalyst 9400 Supervisor Engine-1XL-Y`, etc.
	SupervisorProductNameOrdinal *float64                                                                                           `json:"supervisorProductNameOrdinal,omitempty"` // Supervisor Engine Module Ordinal, supported by the `productNameOrdinal`. Example: The `286315691` chassis ordinal is capable of supporting  different supervisor engine module ordinals: `286316172`, `286316710`, `286320394` etc.
	NetworkDeviceCount           *int                                                                                               `json:"networkDeviceCount,omitempty"`           // Count of network devices
	ImageSummary                 *ResponseSoftwareImageManagementSwimReturnsNetworkDeviceProductNamesForASiteV1ResponseImageSummary `json:"imageSummary,omitempty"`                 //
}
type ResponseSoftwareImageManagementSwimReturnsNetworkDeviceProductNamesForASiteV1ResponseImageSummary struct {
	InstalledImageCount        *int `json:"installedImageCount,omitempty"`        // Count of installed images
	GoldenImageCount           *int `json:"goldenImageCount,omitempty"`           // Count of golden tagged images
	InstalledImageAdvisorCount *int `json:"installedImageAdvisorCount,omitempty"` // Count of advisor on installed images
}
type ResponseSoftwareImageManagementSwimReturnsTheCountOfNetworkDeviceProductNamesForASiteV1 struct {
	Response *ResponseSoftwareImageManagementSwimReturnsTheCountOfNetworkDeviceProductNamesForASiteV1Response `json:"response,omitempty"` //
	Version  string                                                                                           `json:"version,omitempty"`  // The version of the response
}
type ResponseSoftwareImageManagementSwimReturnsTheCountOfNetworkDeviceProductNamesForASiteV1Response struct {
	Count *int `json:"count,omitempty"` // Reports a count, for example, a total count of records in a given resource.
}
type RequestSoftwareImageManagementSwimTriggerSoftwareImageActivationV1 []RequestItemSoftwareImageManagementSwimTriggerSoftwareImageActivationV1 // Array of RequestSoftwareImageManagementSwimTriggerSoftwareImageActivationV1
type RequestItemSoftwareImageManagementSwimTriggerSoftwareImageActivationV1 struct {
	ActivateLowerImageVersion *bool    `json:"activateLowerImageVersion,omitempty"` //
	DeviceUpgradeMode         string   `json:"deviceUpgradeMode,omitempty"`         //
	DeviceUUID                string   `json:"deviceUuid,omitempty"`                //
	DistributeIfNeeded        *bool    `json:"distributeIfNeeded,omitempty"`        //
	ImageUUIDList             []string `json:"imageUuidList,omitempty"`             //
	SmuImageUUIDList          []string `json:"smuImageUuidList,omitempty"`          //
}
type RequestSoftwareImageManagementSwimTriggerSoftwareImageDistributionV1 []RequestItemSoftwareImageManagementSwimTriggerSoftwareImageDistributionV1 // Array of RequestSoftwareImageManagementSwimTriggerSoftwareImageDistributionV1
type RequestItemSoftwareImageManagementSwimTriggerSoftwareImageDistributionV1 struct {
	DeviceUUID string `json:"deviceUuid,omitempty"` //
	ImageUUID  string `json:"imageUuid,omitempty"`  //
}
type RequestSoftwareImageManagementSwimTagAsGoldenImageV1 struct {
	ImageID                string `json:"imageId,omitempty"`                // imageId in uuid format.
	SiteID                 string `json:"siteId,omitempty"`                 // SiteId in uuid format. For Global Site "-1" to be used.
	DeviceRole             string `json:"deviceRole,omitempty"`             // Device Role. Permissible Values : ALL, UNKNOWN, ACCESS, BORDER ROUTER, DISTRIBUTION and CORE.
	DeviceFamilyIDentifier string `json:"deviceFamilyIdentifier,omitempty"` // Device Family Identifier e.g. : 277696480-283933147, 277696480
}
type RequestSoftwareImageManagementSwimImportSoftwareImageViaURLV1 []RequestItemSoftwareImageManagementSwimImportSoftwareImageViaURLV1 // Array of RequestSoftwareImageManagementSwimImportSoftwareImageViaURLV1
type RequestItemSoftwareImageManagementSwimImportSoftwareImageViaURLV1 struct {
	ApplicationType string `json:"applicationType,omitempty"` //
	ImageFamily     string `json:"imageFamily,omitempty"`     //
	SourceURL       string `json:"sourceURL,omitempty"`       //
	ThirdParty      *bool  `json:"thirdParty,omitempty"`      //
	Vendor          string `json:"vendor,omitempty"`          //
}
type RequestSoftwareImageManagementSwimAddImageDistributionServerV1 struct {
	ServerAddress string   `json:"serverAddress,omitempty"` // FQDN or IP address of the server
	Username      string   `json:"username,omitempty"`      // Server username
	PortNumber    *float64 `json:"portNumber,omitempty"`    // Port number
	RootLocation  string   `json:"rootLocation,omitempty"`  // Server root location
	Password      string   `json:"password,omitempty"`      // Server password
}
type RequestSoftwareImageManagementSwimUpdateRemoteImageDistributionServerV1 struct {
	Username   string   `json:"username,omitempty"`   // Server username
	PortNumber *float64 `json:"portNumber,omitempty"` // Port number
	Password   string   `json:"password,omitempty"`   // Server password
}
type RequestSoftwareImageManagementSwimAssignNetworkDeviceProductNameToTheGivenSoftwareImageV1 struct {
	ProductNameOrdinal *float64 `json:"productNameOrdinal,omitempty"` // Product name ordinal is unique value for each network device product
	SiteIDs            []string `json:"siteIds,omitempty"`            // Sites where this image needs to be assigned. Ref https://developer.cisco.com/docs/dna-center/#!sites
}
type RequestSoftwareImageManagementSwimUpdateTheListOfSitesForTheNetworkDeviceProductNameAssignedToTheSoftwareImageV1 struct {
	SiteIDs []string `json:"siteIds,omitempty"` // Sites where all this image need to be assigned
}

//GetSoftwareImageDetailsV1 Get software image details - 0c8f-7a0b-49b9-aedd
/* Returns software image list based on a filter criteria. For example: "filterbyName = cat3k%"


@param GetSoftwareImageDetailsV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-software-image-details-v1
*/
func (s *SoftwareImageManagementSwimService) GetSoftwareImageDetailsV1(GetSoftwareImageDetailsV1QueryParams *GetSoftwareImageDetailsV1QueryParams) (*ResponseSoftwareImageManagementSwimGetSoftwareImageDetailsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/image/importation"

	queryString, _ := query.Values(GetSoftwareImageDetailsV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSoftwareImageManagementSwimGetSoftwareImageDetailsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSoftwareImageDetailsV1(GetSoftwareImageDetailsV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetSoftwareImageDetailsV1")
	}

	result := response.Result().(*ResponseSoftwareImageManagementSwimGetSoftwareImageDetailsV1)
	return result, response, err

}

//GetDeviceFamilyIDentifiersV1 Get Device Family Identifiers - 35ae-1bec-4bd8-89fc
/* API to get Device Family Identifiers for all Device Families that can be used for tagging an image golden.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-family-identifiers-v1
*/
func (s *SoftwareImageManagementSwimService) GetDeviceFamilyIDentifiersV1() (*ResponseSoftwareImageManagementSwimGetDeviceFamilyIDentifiersV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/image/importation/device-family-identifiers"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSoftwareImageManagementSwimGetDeviceFamilyIDentifiersV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDeviceFamilyIDentifiersV1()
		}
		return nil, response, fmt.Errorf("error with operation GetDeviceFamilyIdentifiersV1")
	}

	result := response.Result().(*ResponseSoftwareImageManagementSwimGetDeviceFamilyIDentifiersV1)
	return result, response, err

}

//GetGoldenTagStatusOfAnImageV1 Get Golden Tag Status of an Image. - 96bf-b9b4-419b-a6d0
/* Get golden tag status of an image. Set siteId as -1 for Global site.


@param siteID siteId path parameter. Site Id in uuid format. Set siteId as -1 for Global site.

@param deviceFamilyIDentifier deviceFamilyIdentifier path parameter. Device family identifier e.g. : 277696480-283933147, e.g. : 277696480

@param deviceRole deviceRole path parameter. Device Role. Permissible Values : ALL, UNKNOWN, ACCESS, BORDER ROUTER, DISTRIBUTION and CORE.

@param imageID imageId path parameter. Image Id in uuid format.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-golden-tag-status-of-an-image-v1
*/
func (s *SoftwareImageManagementSwimService) GetGoldenTagStatusOfAnImageV1(siteID string, deviceFamilyIDentifier string, deviceRole string, imageID string) (*ResponseSoftwareImageManagementSwimGetGoldenTagStatusOfAnImageV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/image/importation/golden/site/{siteId}/family/{deviceFamilyIdentifier}/role/{deviceRole}/image/{imageId}"
	path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteID), -1)
	path = strings.Replace(path, "{deviceFamilyIdentifier}", fmt.Sprintf("%v", deviceFamilyIDentifier), -1)
	path = strings.Replace(path, "{deviceRole}", fmt.Sprintf("%v", deviceRole), -1)
	path = strings.Replace(path, "{imageId}", fmt.Sprintf("%v", imageID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSoftwareImageManagementSwimGetGoldenTagStatusOfAnImageV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetGoldenTagStatusOfAnImageV1(siteID, deviceFamilyIDentifier, deviceRole, imageID)
		}
		return nil, response, fmt.Errorf("error with operation GetGoldenTagStatusOfAnImageV1")
	}

	result := response.Result().(*ResponseSoftwareImageManagementSwimGetGoldenTagStatusOfAnImageV1)
	return result, response, err

}

//ReturnsListOfSoftwareImagesV1 Returns list of software images - e4a3-6a8c-48fa-91ea
/* A list of available images for the specified site is provided. The default value of the site is set to global. The list includes images that have been imported onto the disk, as well as the latest and suggested images from Cisco.com.


@param ReturnsListOfSoftwareImagesV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!returns-list-of-software-images-v1
*/
func (s *SoftwareImageManagementSwimService) ReturnsListOfSoftwareImagesV1(ReturnsListOfSoftwareImagesV1QueryParams *ReturnsListOfSoftwareImagesV1QueryParams) (*ResponseSoftwareImageManagementSwimReturnsListOfSoftwareImagesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/images"

	queryString, _ := query.Values(ReturnsListOfSoftwareImagesV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSoftwareImageManagementSwimReturnsListOfSoftwareImagesV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.ReturnsListOfSoftwareImagesV1(ReturnsListOfSoftwareImagesV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation ReturnsListOfSoftwareImagesV1")
	}

	result := response.Result().(*ResponseSoftwareImageManagementSwimReturnsListOfSoftwareImagesV1)
	return result, response, err

}

//ReturnsCountOfSoftwareImagesV1 Returns count of software images - 1391-aa45-4098-8eac
/* Returns the count of software images for given `siteId`. The default value of siteId is global


@param ReturnsCountOfSoftwareImagesV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!returns-count-of-software-images-v1
*/
func (s *SoftwareImageManagementSwimService) ReturnsCountOfSoftwareImagesV1(ReturnsCountOfSoftwareImagesV1QueryParams *ReturnsCountOfSoftwareImagesV1QueryParams) (*ResponseSoftwareImageManagementSwimReturnsCountOfSoftwareImagesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/images/count"

	queryString, _ := query.Values(ReturnsCountOfSoftwareImagesV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSoftwareImageManagementSwimReturnsCountOfSoftwareImagesV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.ReturnsCountOfSoftwareImagesV1(ReturnsCountOfSoftwareImagesV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation ReturnsCountOfSoftwareImagesV1")
	}

	result := response.Result().(*ResponseSoftwareImageManagementSwimReturnsCountOfSoftwareImagesV1)
	return result, response, err

}

//RetrieveImageDistributionServersV1 Retrieve image distribution servers - 7982-39ee-4aaa-aa72
/* Retrieve the list of remote image distribution servers. There can be up to two remote servers.Product always acts as local distribution server, and it is not part of this API response.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieve-image-distribution-servers-v1
*/
func (s *SoftwareImageManagementSwimService) RetrieveImageDistributionServersV1() (*ResponseSoftwareImageManagementSwimRetrieveImageDistributionServersV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/images/distributionServerSettings"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSoftwareImageManagementSwimRetrieveImageDistributionServersV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrieveImageDistributionServersV1()
		}
		return nil, response, fmt.Errorf("error with operation RetrieveImageDistributionServersV1")
	}

	result := response.Result().(*ResponseSoftwareImageManagementSwimRetrieveImageDistributionServersV1)
	return result, response, err

}

//RetrieveSpecificImageDistributionServerV1 Retrieve specific image distribution server - b1ac-99fe-47a9-9c85
/* Retrieve image distribution server for the given server identifier


@param id id path parameter. Server identifier


Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieve-specific-image-distribution-server-v1
*/
func (s *SoftwareImageManagementSwimService) RetrieveSpecificImageDistributionServerV1(id string) (*ResponseSoftwareImageManagementSwimRetrieveSpecificImageDistributionServerV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/images/distributionServerSettings/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSoftwareImageManagementSwimRetrieveSpecificImageDistributionServerV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrieveSpecificImageDistributionServerV1(id)
		}
		return nil, response, fmt.Errorf("error with operation RetrieveSpecificImageDistributionServerV1")
	}

	result := response.Result().(*ResponseSoftwareImageManagementSwimRetrieveSpecificImageDistributionServerV1)
	return result, response, err

}

//RetrieveApplicableAddOnImagesForTheGivenSoftwareImageV1 Retrieve applicable add-on images for the given software image - f0ae-9b76-4ee8-ac84
/* Retrieves the list of applicable add-on images if available for the given software image. `id` can be obtained from the response of API [ /dna/intent/api/v1/images?hasAddonImages=true ].


@param id id path parameter. Software image identifier. Check `/dna/intent/api/v1/images?hasAddonImages=true` API to get the same.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieve-applicable-add-on-images-for-the-given-software-image-v1
*/
func (s *SoftwareImageManagementSwimService) RetrieveApplicableAddOnImagesForTheGivenSoftwareImageV1(id string) (*ResponseSoftwareImageManagementSwimRetrieveApplicableAddOnImagesForTheGivenSoftwareImageV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/images/{id}/addonImages"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSoftwareImageManagementSwimRetrieveApplicableAddOnImagesForTheGivenSoftwareImageV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrieveApplicableAddOnImagesForTheGivenSoftwareImageV1(id)
		}
		return nil, response, fmt.Errorf("error with operation RetrieveApplicableAddOnImagesForTheGivenSoftwareImageV1")
	}

	result := response.Result().(*ResponseSoftwareImageManagementSwimRetrieveApplicableAddOnImagesForTheGivenSoftwareImageV1)
	return result, response, err

}

//ReturnsCountOfAddOnImagesV1 Returns count of add-on images - f0ba-68e0-4acb-8234
/* Count of add-on images available for the given software image identifier, `id` can be obtained from the response of API [ /dna/intent/api/v1/images?hasAddonImages=true ].


@param id id path parameter. Software image identifier. Check API `/dna/intent/api/v1/images` for id from response.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!returns-count-of-add-on-images-v1
*/
func (s *SoftwareImageManagementSwimService) ReturnsCountOfAddOnImagesV1(id string) (*ResponseSoftwareImageManagementSwimReturnsCountOfAddOnImagesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/images/{id}/addonImages/count"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSoftwareImageManagementSwimReturnsCountOfAddOnImagesV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.ReturnsCountOfAddOnImagesV1(id)
		}
		return nil, response, fmt.Errorf("error with operation ReturnsCountOfAddOnImagesV1")
	}

	result := response.Result().(*ResponseSoftwareImageManagementSwimReturnsCountOfAddOnImagesV1)
	return result, response, err

}

//RetrievesNetworkDeviceProductNamesAssignedToASoftwareImageV1 Retrieves network device product names assigned to a software image. - 14a1-f9d8-49b9-80be
/* Returns a list of network device product names and associated sites for a given image identifier. Refer `/dna/intent/api/v1/images` API for obtaining `imageId`.


@param imageID imageId path parameter. Software image identifier. Refer `/dna/intent/api/v1/images` API for obtaining `imageId`

@param RetrievesNetworkDeviceProductNamesAssignedToASoftwareImageV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-network-device-product-names-assigned-to-a-software-image-v1
*/
func (s *SoftwareImageManagementSwimService) RetrievesNetworkDeviceProductNamesAssignedToASoftwareImageV1(imageID string, RetrievesNetworkDeviceProductNamesAssignedToASoftwareImageV1QueryParams *RetrievesNetworkDeviceProductNamesAssignedToASoftwareImageV1QueryParams) (*ResponseSoftwareImageManagementSwimRetrievesNetworkDeviceProductNamesAssignedToASoftwareImageV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/images/{imageId}/siteWiseProductNames"
	path = strings.Replace(path, "{imageId}", fmt.Sprintf("%v", imageID), -1)

	queryString, _ := query.Values(RetrievesNetworkDeviceProductNamesAssignedToASoftwareImageV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSoftwareImageManagementSwimRetrievesNetworkDeviceProductNamesAssignedToASoftwareImageV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesNetworkDeviceProductNamesAssignedToASoftwareImageV1(imageID, RetrievesNetworkDeviceProductNamesAssignedToASoftwareImageV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesNetworkDeviceProductNamesAssignedToASoftwareImageV1")
	}

	result := response.Result().(*ResponseSoftwareImageManagementSwimRetrievesNetworkDeviceProductNamesAssignedToASoftwareImageV1)
	return result, response, err

}

//RetrievesTheCountOfAssignedNetworkDeviceProductsV1 Retrieves the count of assigned network device products - e994-0a33-409b-b8be
/* Returns count of assigned network device product for a given image identifier. Refer `/dna/intent/api/v1/images` API for obtaining `imageId`


@param imageID imageId path parameter. Software image identifier. Refer `/dna/intent/api/v/images` API for obtaining `imageId`

@param RetrievesTheCountOfAssignedNetworkDeviceProductsV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-count-of-assigned-network-device-products-v1
*/
func (s *SoftwareImageManagementSwimService) RetrievesTheCountOfAssignedNetworkDeviceProductsV1(imageID string, RetrievesTheCountOfAssignedNetworkDeviceProductsV1QueryParams *RetrievesTheCountOfAssignedNetworkDeviceProductsV1QueryParams) (*ResponseSoftwareImageManagementSwimRetrievesTheCountOfAssignedNetworkDeviceProductsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/images/{imageId}/siteWiseProductNames/count"
	path = strings.Replace(path, "{imageId}", fmt.Sprintf("%v", imageID), -1)

	queryString, _ := query.Values(RetrievesTheCountOfAssignedNetworkDeviceProductsV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSoftwareImageManagementSwimRetrievesTheCountOfAssignedNetworkDeviceProductsV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheCountOfAssignedNetworkDeviceProductsV1(imageID, RetrievesTheCountOfAssignedNetworkDeviceProductsV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesTheCountOfAssignedNetworkDeviceProductsV1")
	}

	result := response.Result().(*ResponseSoftwareImageManagementSwimRetrievesTheCountOfAssignedNetworkDeviceProductsV1)
	return result, response, err

}

//GetNetworkDeviceImageUpdatesV1 Get network device image updates - 75bf-3a9e-467b-af24
/* Returns the list of network device image updates based on the given filter criteria


@param GetNetworkDeviceImageUpdatesV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-device-image-updates-v1
*/
func (s *SoftwareImageManagementSwimService) GetNetworkDeviceImageUpdatesV1(GetNetworkDeviceImageUpdatesV1QueryParams *GetNetworkDeviceImageUpdatesV1QueryParams) (*ResponseSoftwareImageManagementSwimGetNetworkDeviceImageUpdatesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkDeviceImageUpdates"

	queryString, _ := query.Values(GetNetworkDeviceImageUpdatesV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSoftwareImageManagementSwimGetNetworkDeviceImageUpdatesV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetNetworkDeviceImageUpdatesV1(GetNetworkDeviceImageUpdatesV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetNetworkDeviceImageUpdatesV1")
	}

	result := response.Result().(*ResponseSoftwareImageManagementSwimGetNetworkDeviceImageUpdatesV1)
	return result, response, err

}

//CountOfNetworkDeviceImageUpdatesV1 Count of network device image updates - b980-b848-45a8-9987
/* Returns the count of network device image updates based on the given filter criteria


@param CountOfNetworkDeviceImageUpdatesV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!count-of-network-device-image-updates-v1
*/
func (s *SoftwareImageManagementSwimService) CountOfNetworkDeviceImageUpdatesV1(CountOfNetworkDeviceImageUpdatesV1QueryParams *CountOfNetworkDeviceImageUpdatesV1QueryParams) (*ResponseSoftwareImageManagementSwimCountOfNetworkDeviceImageUpdatesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkDeviceImageUpdates/count"

	queryString, _ := query.Values(CountOfNetworkDeviceImageUpdatesV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSoftwareImageManagementSwimCountOfNetworkDeviceImageUpdatesV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.CountOfNetworkDeviceImageUpdatesV1(CountOfNetworkDeviceImageUpdatesV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation CountOfNetworkDeviceImageUpdatesV1")
	}

	result := response.Result().(*ResponseSoftwareImageManagementSwimCountOfNetworkDeviceImageUpdatesV1)
	return result, response, err

}

//RetrievesTheListOfNetworkDeviceProductNamesV1 Retrieves the list of network device product names - a7bf-3baf-4c29-b1ca
/* Get the list of network device product names, their ordinal, and the support PIDs based on filter criteria.


@param RetrievesTheListOfNetworkDeviceProductNamesV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-list-of-network-device-product-names-v1
*/
func (s *SoftwareImageManagementSwimService) RetrievesTheListOfNetworkDeviceProductNamesV1(RetrievesTheListOfNetworkDeviceProductNamesV1QueryParams *RetrievesTheListOfNetworkDeviceProductNamesV1QueryParams) (*ResponseSoftwareImageManagementSwimRetrievesTheListOfNetworkDeviceProductNamesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/productNames"

	queryString, _ := query.Values(RetrievesTheListOfNetworkDeviceProductNamesV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSoftwareImageManagementSwimRetrievesTheListOfNetworkDeviceProductNamesV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheListOfNetworkDeviceProductNamesV1(RetrievesTheListOfNetworkDeviceProductNamesV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesTheListOfNetworkDeviceProductNamesV1")
	}

	result := response.Result().(*ResponseSoftwareImageManagementSwimRetrievesTheListOfNetworkDeviceProductNamesV1)
	return result, response, err

}

//CountOfNetworkProductNamesV1 Count of network product names - baa2-9b3d-45bb-a870
/* Count of product names based on filter criteria


@param CountOfNetworkProductNamesV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!count-of-network-product-names-v1
*/
func (s *SoftwareImageManagementSwimService) CountOfNetworkProductNamesV1(CountOfNetworkProductNamesV1QueryParams *CountOfNetworkProductNamesV1QueryParams) (*ResponseSoftwareImageManagementSwimCountOfNetworkProductNamesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/productNames/count"

	queryString, _ := query.Values(CountOfNetworkProductNamesV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSoftwareImageManagementSwimCountOfNetworkProductNamesV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.CountOfNetworkProductNamesV1(CountOfNetworkProductNamesV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation CountOfNetworkProductNamesV1")
	}

	result := response.Result().(*ResponseSoftwareImageManagementSwimCountOfNetworkProductNamesV1)
	return result, response, err

}

//RetrieveNetworkDeviceProductNameV1 Retrieve network device product name - 3aa8-fb90-4288-b606
/* Get the network device product name, its ordinal, and supported PIDs.


@param productNameOrdinal productNameOrdinal path parameter. Product name ordinal is unique value for each network device product.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieve-network-device-product-name-v1
*/
func (s *SoftwareImageManagementSwimService) RetrieveNetworkDeviceProductNameV1(productNameOrdinal float64) (*ResponseSoftwareImageManagementSwimRetrieveNetworkDeviceProductNameV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/productNames/{productNameOrdinal}"
	path = strings.Replace(path, "{productNameOrdinal}", fmt.Sprintf("%v", productNameOrdinal), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSoftwareImageManagementSwimRetrieveNetworkDeviceProductNameV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrieveNetworkDeviceProductNameV1(productNameOrdinal)
		}
		return nil, response, fmt.Errorf("error with operation RetrieveNetworkDeviceProductNameV1")
	}

	result := response.Result().(*ResponseSoftwareImageManagementSwimRetrieveNetworkDeviceProductNameV1)
	return result, response, err

}

//ReturnsNetworkDeviceProductNamesForASiteV1 Returns network device product names for a site - 20b5-5b0c-4518-9a03
/* Provides network device product names for a site. The default value of `siteId` is global. The response will include the network device count and image summary.


@param ReturnsNetworkDeviceProductNamesForASiteV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!returns-network-device-product-names-for-a-site-v1
*/
func (s *SoftwareImageManagementSwimService) ReturnsNetworkDeviceProductNamesForASiteV1(ReturnsNetworkDeviceProductNamesForASiteV1QueryParams *ReturnsNetworkDeviceProductNamesForASiteV1QueryParams) (*ResponseSoftwareImageManagementSwimReturnsNetworkDeviceProductNamesForASiteV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/siteWiseProductNames"

	queryString, _ := query.Values(ReturnsNetworkDeviceProductNamesForASiteV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSoftwareImageManagementSwimReturnsNetworkDeviceProductNamesForASiteV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.ReturnsNetworkDeviceProductNamesForASiteV1(ReturnsNetworkDeviceProductNamesForASiteV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation ReturnsNetworkDeviceProductNamesForASiteV1")
	}

	result := response.Result().(*ResponseSoftwareImageManagementSwimReturnsNetworkDeviceProductNamesForASiteV1)
	return result, response, err

}

//ReturnsTheCountOfNetworkDeviceProductNamesForASiteV1 Returns the count of network device product names for a site - 018d-a93c-4e4b-8436
/* Returns the count of network device product names for given filters. The default value of `siteId` is global.


@param ReturnsTheCountOfNetworkDeviceProductNamesForASiteV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!returns-the-count-of-network-device-product-names-for-a-site-v1
*/
func (s *SoftwareImageManagementSwimService) ReturnsTheCountOfNetworkDeviceProductNamesForASiteV1(ReturnsTheCountOfNetworkDeviceProductNamesForASiteV1QueryParams *ReturnsTheCountOfNetworkDeviceProductNamesForASiteV1QueryParams) (*ResponseSoftwareImageManagementSwimReturnsTheCountOfNetworkDeviceProductNamesForASiteV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/siteWiseProductNames/count"

	queryString, _ := query.Values(ReturnsTheCountOfNetworkDeviceProductNamesForASiteV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSoftwareImageManagementSwimReturnsTheCountOfNetworkDeviceProductNamesForASiteV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.ReturnsTheCountOfNetworkDeviceProductNamesForASiteV1(ReturnsTheCountOfNetworkDeviceProductNamesForASiteV1QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation ReturnsTheCountOfNetworkDeviceProductNamesForASiteV1")
	}

	result := response.Result().(*ResponseSoftwareImageManagementSwimReturnsTheCountOfNetworkDeviceProductNamesForASiteV1)
	return result, response, err

}

//TriggerSoftwareImageActivationV1 Trigger software image activation - fb9b-eb66-4f2a-ba4c
/* Activates a software image on a given device. Software image must be present in the device flash


@param TriggerSoftwareImageActivationV1HeaderParams Custom header parameters
@param TriggerSoftwareImageActivationV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!trigger-software-image-activation-v1
*/
func (s *SoftwareImageManagementSwimService) TriggerSoftwareImageActivationV1(requestSoftwareImageManagementSwimTriggerSoftwareImageActivationV1 *RequestSoftwareImageManagementSwimTriggerSoftwareImageActivationV1, TriggerSoftwareImageActivationV1HeaderParams *TriggerSoftwareImageActivationV1HeaderParams, TriggerSoftwareImageActivationV1QueryParams *TriggerSoftwareImageActivationV1QueryParams) (*ResponseSoftwareImageManagementSwimTriggerSoftwareImageActivationV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/image/activation/device"

	queryString, _ := query.Values(TriggerSoftwareImageActivationV1QueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if TriggerSoftwareImageActivationV1HeaderParams != nil {

		if TriggerSoftwareImageActivationV1HeaderParams.ClientType != "" {
			clientRequest = clientRequest.SetHeader("Client-Type", TriggerSoftwareImageActivationV1HeaderParams.ClientType)
		}

		if TriggerSoftwareImageActivationV1HeaderParams.ClientURL != "" {
			clientRequest = clientRequest.SetHeader("Client-Url", TriggerSoftwareImageActivationV1HeaderParams.ClientURL)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetBody(requestSoftwareImageManagementSwimTriggerSoftwareImageActivationV1).
		SetResult(&ResponseSoftwareImageManagementSwimTriggerSoftwareImageActivationV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.TriggerSoftwareImageActivationV1(requestSoftwareImageManagementSwimTriggerSoftwareImageActivationV1, TriggerSoftwareImageActivationV1HeaderParams, TriggerSoftwareImageActivationV1QueryParams)
		}

		return nil, response, fmt.Errorf("error with operation TriggerSoftwareImageActivationV1")
	}

	result := response.Result().(*ResponseSoftwareImageManagementSwimTriggerSoftwareImageActivationV1)
	return result, response, err

}

//TriggerSoftwareImageDistributionV1 Trigger software image distribution - 8cb6-783b-4fab-a1f4
/* Distributes a software image on a given device. Software image must be imported successfully into DNA Center before it can be distributed



Documentation Link: https://developer.cisco.com/docs/dna-center/#!trigger-software-image-distribution-v1
*/
func (s *SoftwareImageManagementSwimService) TriggerSoftwareImageDistributionV1(requestSoftwareImageManagementSwimTriggerSoftwareImageDistributionV1 *RequestSoftwareImageManagementSwimTriggerSoftwareImageDistributionV1) (*ResponseSoftwareImageManagementSwimTriggerSoftwareImageDistributionV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/image/distribution"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSoftwareImageManagementSwimTriggerSoftwareImageDistributionV1).
		SetResult(&ResponseSoftwareImageManagementSwimTriggerSoftwareImageDistributionV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.TriggerSoftwareImageDistributionV1(requestSoftwareImageManagementSwimTriggerSoftwareImageDistributionV1)
		}

		return nil, response, fmt.Errorf("error with operation TriggerSoftwareImageDistributionV1")
	}

	result := response.Result().(*ResponseSoftwareImageManagementSwimTriggerSoftwareImageDistributionV1)
	return result, response, err

}

//TagAsGoldenImageV1 Tag as Golden Image - 5c91-7a67-474b-a0e0
/* Golden Tag image. Set siteId as -1 for Global site.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!tag-as-golden-image-v1
*/
func (s *SoftwareImageManagementSwimService) TagAsGoldenImageV1(requestSoftwareImageManagementSwimTagAsGoldenImageV1 *RequestSoftwareImageManagementSwimTagAsGoldenImageV1) (*ResponseSoftwareImageManagementSwimTagAsGoldenImageV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/image/importation/golden"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSoftwareImageManagementSwimTagAsGoldenImageV1).
		SetResult(&ResponseSoftwareImageManagementSwimTagAsGoldenImageV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.TagAsGoldenImageV1(requestSoftwareImageManagementSwimTagAsGoldenImageV1)
		}

		return nil, response, fmt.Errorf("error with operation TagAsGoldenImageV1")
	}

	result := response.Result().(*ResponseSoftwareImageManagementSwimTagAsGoldenImageV1)
	return result, response, err

}

//ImportLocalSoftwareImageV1 Import local software image - 1491-f90f-48da-aabe
/* Fetches a software image from local file system and uploads to DNA Center. Supported software image files extensions are bin, img, tar, smu, pie, aes, iso, ova, tar_gz and qcow2


@param ImportLocalSoftwareImageV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!import-local-software-image-v1
*/
func (s *SoftwareImageManagementSwimService) ImportLocalSoftwareImageV1(ImportLocalSoftwareImageV1QueryParams *ImportLocalSoftwareImageV1QueryParams, ImportLocalSoftwareImageMultipartFields *ImportLocalSoftwareImageMultipartFields) (*ResponseSoftwareImageManagementSwimImportLocalSoftwareImageV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/image/importation/source/file"

	queryString, _ := query.Values(ImportLocalSoftwareImageV1QueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if ImportLocalSoftwareImageMultipartFields != nil {
		clientRequest = clientRequest.SetFileReader("file", ImportLocalSoftwareImageMultipartFields.FileName, ImportLocalSoftwareImageMultipartFields.File)
	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).
		SetResult(&ResponseSoftwareImageManagementSwimImportLocalSoftwareImageV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.ImportLocalSoftwareImageV1(ImportLocalSoftwareImageV1QueryParams, ImportLocalSoftwareImageMultipartFields)
		}

		return nil, response, fmt.Errorf("error with operation ImportLocalSoftwareImageV1")
	}

	result := response.Result().(*ResponseSoftwareImageManagementSwimImportLocalSoftwareImageV1)
	return result, response, err

}

//ImportSoftwareImageViaURLV1 Import software image via URL - bc8a-ab47-46ca-883d
/* Fetches a software image from remote file system (using URL for HTTP/FTP) and uploads to DNA Center. Supported image files extensions are bin, img, tar, smu, pie, aes, iso, ova, tar_gz and qcow2


@param ImportSoftwareImageViaURLV1QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!import-software-image-via-url-v1
*/
func (s *SoftwareImageManagementSwimService) ImportSoftwareImageViaURLV1(requestSoftwareImageManagementSwimImportSoftwareImageViaURLV1 *RequestSoftwareImageManagementSwimImportSoftwareImageViaURLV1, ImportSoftwareImageViaURLV1QueryParams *ImportSoftwareImageViaURLV1QueryParams) (*ResponseSoftwareImageManagementSwimImportSoftwareImageViaURLV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/image/importation/source/url"

	queryString, _ := query.Values(ImportSoftwareImageViaURLV1QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetBody(requestSoftwareImageManagementSwimImportSoftwareImageViaURLV1).
		SetResult(&ResponseSoftwareImageManagementSwimImportSoftwareImageViaURLV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.ImportSoftwareImageViaURLV1(requestSoftwareImageManagementSwimImportSoftwareImageViaURLV1, ImportSoftwareImageViaURLV1QueryParams)
		}

		return nil, response, fmt.Errorf("error with operation ImportSoftwareImageViaUrlV1")
	}

	result := response.Result().(*ResponseSoftwareImageManagementSwimImportSoftwareImageViaURLV1)
	return result, response, err

}

//AddImageDistributionServerV1 Add image distribution server - 0699-0a83-4aaa-be15
/* Add remote server for distributing software images. Upto two such distribution servers are supported.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-image-distribution-server-v1
*/
func (s *SoftwareImageManagementSwimService) AddImageDistributionServerV1(requestSoftwareImageManagementSwimAddImageDistributionServerV1 *RequestSoftwareImageManagementSwimAddImageDistributionServerV1) (*ResponseSoftwareImageManagementSwimAddImageDistributionServerV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/images/distributionServerSettings"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSoftwareImageManagementSwimAddImageDistributionServerV1).
		SetResult(&ResponseSoftwareImageManagementSwimAddImageDistributionServerV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddImageDistributionServerV1(requestSoftwareImageManagementSwimAddImageDistributionServerV1)
		}

		return nil, response, fmt.Errorf("error with operation AddImageDistributionServerV1")
	}

	result := response.Result().(*ResponseSoftwareImageManagementSwimAddImageDistributionServerV1)
	return result, response, err

}

//DownloadTheSoftwareImageV1 Download the software image - c382-2b30-447a-a189
/* Initiates download of the software image from Cisco.com on the disk for the given `id`. Refer to `/dna/intent/api/v1/images` for obtaining `id`.


@param id id path parameter. Software image identifier. Check API `/dna/intent/api/v1/images` for `id` from response.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!download-the-software-image-v1
*/
func (s *SoftwareImageManagementSwimService) DownloadTheSoftwareImageV1(id string) (*ResponseSoftwareImageManagementSwimDownloadTheSoftwareImageV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/images/{id}/download"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSoftwareImageManagementSwimDownloadTheSoftwareImageV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.DownloadTheSoftwareImageV1(id)
		}

		return nil, response, fmt.Errorf("error with operation DownloadTheSoftwareImageV1")
	}

	result := response.Result().(*ResponseSoftwareImageManagementSwimDownloadTheSoftwareImageV1)
	return result, response, err

}

//AssignNetworkDeviceProductNameToTheGivenSoftwareImageV1 Assign network device product name to the given software image - 0089-283d-4609-98a5
/* Assign network device product name and sites for the given image identifier. Refer `/dna/intent/api/v1/images` API for obtaining imageId


@param imageID imageId path parameter. Software image identifier. Refer `/dna/intent/api/v1/images` API for obtaining `imageId`


Documentation Link: https://developer.cisco.com/docs/dna-center/#!assign-network-device-product-name-to-the-given-software-image-v1
*/
func (s *SoftwareImageManagementSwimService) AssignNetworkDeviceProductNameToTheGivenSoftwareImageV1(imageID string, requestSoftwareImageManagementSwimAssignNetworkDeviceProductNameToTheGivenSoftwareImageV1 *RequestSoftwareImageManagementSwimAssignNetworkDeviceProductNameToTheGivenSoftwareImageV1) (*ResponseSoftwareImageManagementSwimAssignNetworkDeviceProductNameToTheGivenSoftwareImageV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/images/{imageId}/siteWiseProductNames"
	path = strings.Replace(path, "{imageId}", fmt.Sprintf("%v", imageID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSoftwareImageManagementSwimAssignNetworkDeviceProductNameToTheGivenSoftwareImageV1).
		SetResult(&ResponseSoftwareImageManagementSwimAssignNetworkDeviceProductNameToTheGivenSoftwareImageV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AssignNetworkDeviceProductNameToTheGivenSoftwareImageV1(imageID, requestSoftwareImageManagementSwimAssignNetworkDeviceProductNameToTheGivenSoftwareImageV1)
		}

		return nil, response, fmt.Errorf("error with operation AssignNetworkDeviceProductNameToTheGivenSoftwareImageV1")
	}

	result := response.Result().(*ResponseSoftwareImageManagementSwimAssignNetworkDeviceProductNameToTheGivenSoftwareImageV1)
	return result, response, err

}

//UpdateRemoteImageDistributionServerV1 Update remote image distribution server - 2caa-c9cc-469b-a3d5
/* Update remote image distribution server details.


@param id id path parameter. Remote server identifier.

*/
func (s *SoftwareImageManagementSwimService) UpdateRemoteImageDistributionServerV1(id string, requestSoftwareImageManagementSwimUpdateRemoteImageDistributionServerV1 *RequestSoftwareImageManagementSwimUpdateRemoteImageDistributionServerV1) (*ResponseSoftwareImageManagementSwimUpdateRemoteImageDistributionServerV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/images/distributionServerSettings/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSoftwareImageManagementSwimUpdateRemoteImageDistributionServerV1).
		SetResult(&ResponseSoftwareImageManagementSwimUpdateRemoteImageDistributionServerV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateRemoteImageDistributionServerV1(id, requestSoftwareImageManagementSwimUpdateRemoteImageDistributionServerV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdateRemoteImageDistributionServerV1")
	}

	result := response.Result().(*ResponseSoftwareImageManagementSwimUpdateRemoteImageDistributionServerV1)
	return result, response, err

}

//UpdateTheListOfSitesForTheNetworkDeviceProductNameAssignedToTheSoftwareImageV1 Update the list of sites for the network device product name assigned to the software image - 1da6-d80b-40fa-8bdc
/* Update the list of sites for the network device product name assigned to the software image. Refer to `/dna/intent/api/v1/images` and `/dna/intent/api/v1/images/{imageId}/siteWiseProductNames` GET APIs for obtaining  `imageId` and `productNameOrdinal` respectively.


@param imageID imageId path parameter. Software image identifier. Refer `/dna/intent/api/v1/images` API for obtaining `imageId`

@param productNameOrdinal productNameOrdinal path parameter. Product name ordinal is unique value for each network device product. Refer `/dna/intent/api/v1/images/{imageId}/siteWiseProductNames` GET API for obtaining `productNameOrdinal`.

*/
func (s *SoftwareImageManagementSwimService) UpdateTheListOfSitesForTheNetworkDeviceProductNameAssignedToTheSoftwareImageV1(imageID string, productNameOrdinal float64, requestSoftwareImageManagementSwimUpdateTheListOfSitesForTheNetworkDeviceProductNameAssignedToTheSoftwareImageV1 *RequestSoftwareImageManagementSwimUpdateTheListOfSitesForTheNetworkDeviceProductNameAssignedToTheSoftwareImageV1) (*ResponseSoftwareImageManagementSwimUpdateTheListOfSitesForTheNetworkDeviceProductNameAssignedToTheSoftwareImageV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/images/{imageId}/siteWiseProductNames/{productNameOrdinal}"
	path = strings.Replace(path, "{imageId}", fmt.Sprintf("%v", imageID), -1)
	path = strings.Replace(path, "{productNameOrdinal}", fmt.Sprintf("%v", productNameOrdinal), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSoftwareImageManagementSwimUpdateTheListOfSitesForTheNetworkDeviceProductNameAssignedToTheSoftwareImageV1).
		SetResult(&ResponseSoftwareImageManagementSwimUpdateTheListOfSitesForTheNetworkDeviceProductNameAssignedToTheSoftwareImageV1{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateTheListOfSitesForTheNetworkDeviceProductNameAssignedToTheSoftwareImageV1(imageID, productNameOrdinal, requestSoftwareImageManagementSwimUpdateTheListOfSitesForTheNetworkDeviceProductNameAssignedToTheSoftwareImageV1)
		}
		return nil, response, fmt.Errorf("error with operation UpdateTheListOfSitesForTheNetworkDeviceProductNameAssignedToTheSoftwareImageV1")
	}

	result := response.Result().(*ResponseSoftwareImageManagementSwimUpdateTheListOfSitesForTheNetworkDeviceProductNameAssignedToTheSoftwareImageV1)
	return result, response, err

}

//RemoveGoldenTagForImageV1 Remove Golden Tag for image. - f3b9-5978-4f6a-897a
/* Remove golden tag. Set siteId as -1 for Global site.


@param siteID siteId path parameter. Site Id in uuid format. Set siteId as -1 for Global site.

@param deviceFamilyIDentifier deviceFamilyIdentifier path parameter. Device family identifier e.g. : 277696480-283933147, e.g. : 277696480

@param deviceRole deviceRole path parameter. Device Role. Permissible Values : ALL, UNKNOWN, ACCESS, BORDER ROUTER, DISTRIBUTION and CORE.

@param imageID imageId path parameter. Image Id in uuid format.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!remove-golden-tag-for-image-v1
*/
func (s *SoftwareImageManagementSwimService) RemoveGoldenTagForImageV1(siteID string, deviceFamilyIDentifier string, deviceRole string, imageID string) (*ResponseSoftwareImageManagementSwimRemoveGoldenTagForImageV1, *resty.Response, error) {
	//siteID string,deviceFamilyIDentifier string,deviceRole string,imageID string
	path := "/dna/intent/api/v1/image/importation/golden/site/{siteId}/family/{deviceFamilyIdentifier}/role/{deviceRole}/image/{imageId}"
	path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteID), -1)
	path = strings.Replace(path, "{deviceFamilyIdentifier}", fmt.Sprintf("%v", deviceFamilyIDentifier), -1)
	path = strings.Replace(path, "{deviceRole}", fmt.Sprintf("%v", deviceRole), -1)
	path = strings.Replace(path, "{imageId}", fmt.Sprintf("%v", imageID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSoftwareImageManagementSwimRemoveGoldenTagForImageV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RemoveGoldenTagForImageV1(siteID, deviceFamilyIDentifier, deviceRole, imageID)
		}
		return nil, response, fmt.Errorf("error with operation RemoveGoldenTagForImageV1")
	}

	result := response.Result().(*ResponseSoftwareImageManagementSwimRemoveGoldenTagForImageV1)
	return result, response, err

}

//RemoveImageDistributionServerV1 Remove image distribution server - 43a7-5b17-404a-aa9f
/* Delete remote image distribution server.


@param id id path parameter. Remote server identifier.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!remove-image-distribution-server-v1
*/
func (s *SoftwareImageManagementSwimService) RemoveImageDistributionServerV1(id string) (*ResponseSoftwareImageManagementSwimRemoveImageDistributionServerV1, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/images/distributionServerSettings/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSoftwareImageManagementSwimRemoveImageDistributionServerV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RemoveImageDistributionServerV1(id)
		}
		return nil, response, fmt.Errorf("error with operation RemoveImageDistributionServerV1")
	}

	result := response.Result().(*ResponseSoftwareImageManagementSwimRemoveImageDistributionServerV1)
	return result, response, err

}

//UnassignNetworkDeviceProductNameFromTheGivenSoftwareImageV1 Unassign network device product name from the given software image - 3fa4-39e3-4a4b-8eaf
/* This API unassigns the network device product name from all the sites for the given software image.
        Refer to `/dna/intent/api/v1/images` and `/dna/intent/api/v1/images/{imageId}/siteWiseProductNames` GET APIs for obtaining  `imageId` and `productNameOrdinal` respectively.


@param imageID imageId path parameter. Software image identifier. Refer `/dna/intent/api/v1/images` API for obtaining `imageId`

@param productNameOrdinal productNameOrdinal path parameter. The product name ordinal is a unique value for each network device product. Refer `/dna/intent/api/v1/images/{imageId}/siteWiseProductNames` GET API for obtaining `productNameOrdinal`


Documentation Link: https://developer.cisco.com/docs/dna-center/#!unassign-network-device-product-name-from-the-given-software-image-v1
*/
func (s *SoftwareImageManagementSwimService) UnassignNetworkDeviceProductNameFromTheGivenSoftwareImageV1(imageID string, productNameOrdinal float64) (*ResponseSoftwareImageManagementSwimUnassignNetworkDeviceProductNameFromTheGivenSoftwareImageV1, *resty.Response, error) {
	//imageID string,productNameOrdinal float64
	path := "/dna/intent/api/v1/images/{imageId}/siteWiseProductNames/{productNameOrdinal}"
	path = strings.Replace(path, "{imageId}", fmt.Sprintf("%v", imageID), -1)
	path = strings.Replace(path, "{productNameOrdinal}", fmt.Sprintf("%v", productNameOrdinal), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSoftwareImageManagementSwimUnassignNetworkDeviceProductNameFromTheGivenSoftwareImageV1{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UnassignNetworkDeviceProductNameFromTheGivenSoftwareImageV1(imageID, productNameOrdinal)
		}
		return nil, response, fmt.Errorf("error with operation UnassignNetworkDeviceProductNameFromTheGivenSoftwareImageV1")
	}

	result := response.Result().(*ResponseSoftwareImageManagementSwimUnassignNetworkDeviceProductNameFromTheGivenSoftwareImageV1)
	return result, response, err

}

// Alias Function
func (s *SoftwareImageManagementSwimService) GetDeviceFamilyIDentifiers() (*ResponseSoftwareImageManagementSwimGetDeviceFamilyIDentifiersV1, *resty.Response, error) {
	return s.GetDeviceFamilyIDentifiersV1()
}

// Alias Function
func (s *SoftwareImageManagementSwimService) TriggerSoftwareImageDistribution(requestSoftwareImageManagementSwimTriggerSoftwareImageDistributionV1 *RequestSoftwareImageManagementSwimTriggerSoftwareImageDistributionV1) (*ResponseSoftwareImageManagementSwimTriggerSoftwareImageDistributionV1, *resty.Response, error) {
	return s.TriggerSoftwareImageDistributionV1(requestSoftwareImageManagementSwimTriggerSoftwareImageDistributionV1)
}

// Alias Function
func (s *SoftwareImageManagementSwimService) GetNetworkDeviceImageUpdates(GetNetworkDeviceImageUpdatesV1QueryParams *GetNetworkDeviceImageUpdatesV1QueryParams) (*ResponseSoftwareImageManagementSwimGetNetworkDeviceImageUpdatesV1, *resty.Response, error) {
	return s.GetNetworkDeviceImageUpdatesV1(GetNetworkDeviceImageUpdatesV1QueryParams)
}

// Alias Function
func (s *SoftwareImageManagementSwimService) UpdateRemoteImageDistributionServer(id string, requestSoftwareImageManagementSwimUpdateRemoteImageDistributionServerV1 *RequestSoftwareImageManagementSwimUpdateRemoteImageDistributionServerV1) (*ResponseSoftwareImageManagementSwimUpdateRemoteImageDistributionServerV1, *resty.Response, error) {
	return s.UpdateRemoteImageDistributionServerV1(id, requestSoftwareImageManagementSwimUpdateRemoteImageDistributionServerV1)
}

// Alias Function
func (s *SoftwareImageManagementSwimService) RetrievesTheListOfNetworkDeviceProductNames(RetrievesTheListOfNetworkDeviceProductNamesV1QueryParams *RetrievesTheListOfNetworkDeviceProductNamesV1QueryParams) (*ResponseSoftwareImageManagementSwimRetrievesTheListOfNetworkDeviceProductNamesV1, *resty.Response, error) {
	return s.RetrievesTheListOfNetworkDeviceProductNamesV1(RetrievesTheListOfNetworkDeviceProductNamesV1QueryParams)
}

// Alias Function
func (s *SoftwareImageManagementSwimService) DownloadTheSoftwareImage(id string) (*ResponseSoftwareImageManagementSwimDownloadTheSoftwareImageV1, *resty.Response, error) {
	return s.DownloadTheSoftwareImageV1(id)
}

// Alias Function
func (s *SoftwareImageManagementSwimService) TriggerSoftwareImageActivation(requestSoftwareImageManagementSwimTriggerSoftwareImageActivationV1 *RequestSoftwareImageManagementSwimTriggerSoftwareImageActivationV1, TriggerSoftwareImageActivationV1HeaderParams *TriggerSoftwareImageActivationV1HeaderParams, TriggerSoftwareImageActivationV1QueryParams *TriggerSoftwareImageActivationV1QueryParams) (*ResponseSoftwareImageManagementSwimTriggerSoftwareImageActivationV1, *resty.Response, error) {
	return s.TriggerSoftwareImageActivationV1(requestSoftwareImageManagementSwimTriggerSoftwareImageActivationV1, TriggerSoftwareImageActivationV1HeaderParams, TriggerSoftwareImageActivationV1QueryParams)
}

// Alias Function
func (s *SoftwareImageManagementSwimService) AssignNetworkDeviceProductNameToTheGivenSoftwareImage(imageID string, requestSoftwareImageManagementSwimAssignNetworkDeviceProductNameToTheGivenSoftwareImageV1 *RequestSoftwareImageManagementSwimAssignNetworkDeviceProductNameToTheGivenSoftwareImageV1) (*ResponseSoftwareImageManagementSwimAssignNetworkDeviceProductNameToTheGivenSoftwareImageV1, *resty.Response, error) {
	return s.AssignNetworkDeviceProductNameToTheGivenSoftwareImageV1(imageID, requestSoftwareImageManagementSwimAssignNetworkDeviceProductNameToTheGivenSoftwareImageV1)
}

// Alias Function
func (s *SoftwareImageManagementSwimService) RetrieveNetworkDeviceProductName(productNameOrdinal float64) (*ResponseSoftwareImageManagementSwimRetrieveNetworkDeviceProductNameV1, *resty.Response, error) {
	return s.RetrieveNetworkDeviceProductNameV1(productNameOrdinal)
}

// Alias Function
func (s *SoftwareImageManagementSwimService) GetSoftwareImageDetails(GetSoftwareImageDetailsV1QueryParams *GetSoftwareImageDetailsV1QueryParams) (*ResponseSoftwareImageManagementSwimGetSoftwareImageDetailsV1, *resty.Response, error) {
	return s.GetSoftwareImageDetailsV1(GetSoftwareImageDetailsV1QueryParams)
}

// Alias Function
func (s *SoftwareImageManagementSwimService) ReturnsListOfSoftwareImages(ReturnsListOfSoftwareImagesV1QueryParams *ReturnsListOfSoftwareImagesV1QueryParams) (*ResponseSoftwareImageManagementSwimReturnsListOfSoftwareImagesV1, *resty.Response, error) {
	return s.ReturnsListOfSoftwareImagesV1(ReturnsListOfSoftwareImagesV1QueryParams)
}

// Alias Function
func (s *SoftwareImageManagementSwimService) ReturnsCountOfAddOnImages(id string) (*ResponseSoftwareImageManagementSwimReturnsCountOfAddOnImagesV1, *resty.Response, error) {
	return s.ReturnsCountOfAddOnImagesV1(id)
}

// Alias Function
func (s *SoftwareImageManagementSwimService) CountOfNetworkProductNames(CountOfNetworkProductNamesV1QueryParams *CountOfNetworkProductNamesV1QueryParams) (*ResponseSoftwareImageManagementSwimCountOfNetworkProductNamesV1, *resty.Response, error) {
	return s.CountOfNetworkProductNamesV1(CountOfNetworkProductNamesV1QueryParams)
}

// Alias Function
func (s *SoftwareImageManagementSwimService) RetrieveSpecificImageDistributionServer(id string) (*ResponseSoftwareImageManagementSwimRetrieveSpecificImageDistributionServerV1, *resty.Response, error) {
	return s.RetrieveSpecificImageDistributionServerV1(id)
}

// Alias Function
func (s *SoftwareImageManagementSwimService) ImportSoftwareImageViaURL(requestSoftwareImageManagementSwimImportSoftwareImageViaURLV1 *RequestSoftwareImageManagementSwimImportSoftwareImageViaURLV1, ImportSoftwareImageViaURLV1QueryParams *ImportSoftwareImageViaURLV1QueryParams) (*ResponseSoftwareImageManagementSwimImportSoftwareImageViaURLV1, *resty.Response, error) {
	return s.ImportSoftwareImageViaURLV1(requestSoftwareImageManagementSwimImportSoftwareImageViaURLV1, ImportSoftwareImageViaURLV1QueryParams)
}

// Alias Function
func (s *SoftwareImageManagementSwimService) ReturnsNetworkDeviceProductNamesForASite(ReturnsNetworkDeviceProductNamesForASiteV1QueryParams *ReturnsNetworkDeviceProductNamesForASiteV1QueryParams) (*ResponseSoftwareImageManagementSwimReturnsNetworkDeviceProductNamesForASiteV1, *resty.Response, error) {
	return s.ReturnsNetworkDeviceProductNamesForASiteV1(ReturnsNetworkDeviceProductNamesForASiteV1QueryParams)
}

// Alias Function
func (s *SoftwareImageManagementSwimService) ReturnsCountOfSoftwareImages(ReturnsCountOfSoftwareImagesV1QueryParams *ReturnsCountOfSoftwareImagesV1QueryParams) (*ResponseSoftwareImageManagementSwimReturnsCountOfSoftwareImagesV1, *resty.Response, error) {
	return s.ReturnsCountOfSoftwareImagesV1(ReturnsCountOfSoftwareImagesV1QueryParams)
}

// Alias Function
func (s *SoftwareImageManagementSwimService) CountOfNetworkDeviceImageUpdates(CountOfNetworkDeviceImageUpdatesV1QueryParams *CountOfNetworkDeviceImageUpdatesV1QueryParams) (*ResponseSoftwareImageManagementSwimCountOfNetworkDeviceImageUpdatesV1, *resty.Response, error) {
	return s.CountOfNetworkDeviceImageUpdatesV1(CountOfNetworkDeviceImageUpdatesV1QueryParams)
}

// Alias Function
func (s *SoftwareImageManagementSwimService) UnassignNetworkDeviceProductNameFromTheGivenSoftwareImage(imageID string, productNameOrdinal float64) (*ResponseSoftwareImageManagementSwimUnassignNetworkDeviceProductNameFromTheGivenSoftwareImageV1, *resty.Response, error) {
	return s.UnassignNetworkDeviceProductNameFromTheGivenSoftwareImageV1(imageID, productNameOrdinal)
}

// Alias Function
func (s *SoftwareImageManagementSwimService) RemoveImageDistributionServer(id string) (*ResponseSoftwareImageManagementSwimRemoveImageDistributionServerV1, *resty.Response, error) {
	return s.RemoveImageDistributionServerV1(id)
}

// Alias Function
func (s *SoftwareImageManagementSwimService) RetrieveApplicableAddOnImagesForTheGivenSoftwareImage(id string) (*ResponseSoftwareImageManagementSwimRetrieveApplicableAddOnImagesForTheGivenSoftwareImageV1, *resty.Response, error) {
	return s.RetrieveApplicableAddOnImagesForTheGivenSoftwareImageV1(id)
}

// Alias Function
func (s *SoftwareImageManagementSwimService) RemoveGoldenTagForImage(siteID string, deviceFamilyIDentifier string, deviceRole string, imageID string) (*ResponseSoftwareImageManagementSwimRemoveGoldenTagForImageV1, *resty.Response, error) {
	return s.RemoveGoldenTagForImageV1(siteID, deviceFamilyIDentifier, deviceRole, imageID)
}

// Alias Function
func (s *SoftwareImageManagementSwimService) ReturnsTheCountOfNetworkDeviceProductNamesForASite(ReturnsTheCountOfNetworkDeviceProductNamesForASiteV1QueryParams *ReturnsTheCountOfNetworkDeviceProductNamesForASiteV1QueryParams) (*ResponseSoftwareImageManagementSwimReturnsTheCountOfNetworkDeviceProductNamesForASiteV1, *resty.Response, error) {
	return s.ReturnsTheCountOfNetworkDeviceProductNamesForASiteV1(ReturnsTheCountOfNetworkDeviceProductNamesForASiteV1QueryParams)
}

// Alias Function
func (s *SoftwareImageManagementSwimService) RetrievesTheCountOfAssignedNetworkDeviceProducts(imageID string, RetrievesTheCountOfAssignedNetworkDeviceProductsV1QueryParams *RetrievesTheCountOfAssignedNetworkDeviceProductsV1QueryParams) (*ResponseSoftwareImageManagementSwimRetrievesTheCountOfAssignedNetworkDeviceProductsV1, *resty.Response, error) {
	return s.RetrievesTheCountOfAssignedNetworkDeviceProductsV1(imageID, RetrievesTheCountOfAssignedNetworkDeviceProductsV1QueryParams)
}

// Alias Function
func (s *SoftwareImageManagementSwimService) GetGoldenTagStatusOfAnImage(siteID string, deviceFamilyIDentifier string, deviceRole string, imageID string) (*ResponseSoftwareImageManagementSwimGetGoldenTagStatusOfAnImageV1, *resty.Response, error) {
	return s.GetGoldenTagStatusOfAnImageV1(siteID, deviceFamilyIDentifier, deviceRole, imageID)
}

// Alias Function
func (s *SoftwareImageManagementSwimService) AddImageDistributionServer(requestSoftwareImageManagementSwimAddImageDistributionServerV1 *RequestSoftwareImageManagementSwimAddImageDistributionServerV1) (*ResponseSoftwareImageManagementSwimAddImageDistributionServerV1, *resty.Response, error) {
	return s.AddImageDistributionServerV1(requestSoftwareImageManagementSwimAddImageDistributionServerV1)
}

// Alias Function
func (s *SoftwareImageManagementSwimService) UpdateTheListOfSitesForTheNetworkDeviceProductNameAssignedToTheSoftwareImage(imageID string, productNameOrdinal float64, requestSoftwareImageManagementSwimUpdateTheListOfSitesForTheNetworkDeviceProductNameAssignedToTheSoftwareImageV1 *RequestSoftwareImageManagementSwimUpdateTheListOfSitesForTheNetworkDeviceProductNameAssignedToTheSoftwareImageV1) (*ResponseSoftwareImageManagementSwimUpdateTheListOfSitesForTheNetworkDeviceProductNameAssignedToTheSoftwareImageV1, *resty.Response, error) {
	return s.UpdateTheListOfSitesForTheNetworkDeviceProductNameAssignedToTheSoftwareImageV1(imageID, productNameOrdinal, requestSoftwareImageManagementSwimUpdateTheListOfSitesForTheNetworkDeviceProductNameAssignedToTheSoftwareImageV1)
}

// Alias Function
func (s *SoftwareImageManagementSwimService) ImportLocalSoftwareImage(ImportLocalSoftwareImageV1QueryParams *ImportLocalSoftwareImageV1QueryParams, ImportLocalSoftwareImageMultipartFields *ImportLocalSoftwareImageMultipartFields) (*ResponseSoftwareImageManagementSwimImportLocalSoftwareImageV1, *resty.Response, error) {
	return s.ImportLocalSoftwareImageV1(ImportLocalSoftwareImageV1QueryParams, ImportLocalSoftwareImageMultipartFields)
}

// Alias Function
func (s *SoftwareImageManagementSwimService) TagAsGoldenImage(requestSoftwareImageManagementSwimTagAsGoldenImageV1 *RequestSoftwareImageManagementSwimTagAsGoldenImageV1) (*ResponseSoftwareImageManagementSwimTagAsGoldenImageV1, *resty.Response, error) {
	return s.TagAsGoldenImageV1(requestSoftwareImageManagementSwimTagAsGoldenImageV1)
}

// Alias Function
func (s *SoftwareImageManagementSwimService) RetrieveImageDistributionServers() (*ResponseSoftwareImageManagementSwimRetrieveImageDistributionServersV1, *resty.Response, error) {
	return s.RetrieveImageDistributionServersV1()
}

// Alias Function
func (s *SoftwareImageManagementSwimService) RetrievesNetworkDeviceProductNamesAssignedToASoftwareImage(imageID string, RetrievesNetworkDeviceProductNamesAssignedToASoftwareImageV1QueryParams *RetrievesNetworkDeviceProductNamesAssignedToASoftwareImageV1QueryParams) (*ResponseSoftwareImageManagementSwimRetrievesNetworkDeviceProductNamesAssignedToASoftwareImageV1, *resty.Response, error) {
	return s.RetrievesNetworkDeviceProductNamesAssignedToASoftwareImageV1(imageID, RetrievesNetworkDeviceProductNamesAssignedToASoftwareImageV1QueryParams)
}
