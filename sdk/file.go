package catalyst

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
)

type FileService service

type UploadFileMultipartFields struct {
	File     io.Reader
	FileName string
}

type ResponseFileGetListOfAvailableNamespacesV1 struct {
	Response []string `json:"response,omitempty"` //
	Version  string   `json:"version,omitempty"`  //
}
type ResponseFileGetListOfFilesV1 struct {
	Response *[]ResponseFileGetListOfFilesV1Response `json:"response,omitempty"` //
	Version  string                                  `json:"version,omitempty"`  //
}
type ResponseFileGetListOfFilesV1Response struct {
	AttributeInfo  *ResponseFileGetListOfFilesV1ResponseAttributeInfo    `json:"attributeInfo,omitempty"`  //
	DownloadPath   string                                                `json:"downloadPath,omitempty"`   //
	Encrypted      *bool                                                 `json:"encrypted,omitempty"`      //
	FileFormat     string                                                `json:"fileFormat,omitempty"`     //
	FileSize       string                                                `json:"fileSize,omitempty"`       //
	ID             string                                                `json:"id,omitempty"`             //
	Md5Checksum    string                                                `json:"md5Checksum,omitempty"`    //
	Name           string                                                `json:"name,omitempty"`           //
	NameSpace      string                                                `json:"nameSpace,omitempty"`      //
	SftpServerList *[]ResponseFileGetListOfFilesV1ResponseSftpServerList `json:"sftpServerList,omitempty"` //
	Sha1Checksum   string                                                `json:"sha1Checksum,omitempty"`   //
	TaskID         string                                                `json:"taskId,omitempty"`         //
}
type ResponseFileGetListOfFilesV1ResponseAttributeInfo interface{}
type ResponseFileGetListOfFilesV1ResponseSftpServerList interface{}
type ResponseFileDownloadAFileByFileIDV1 interface{}
type ResponseFileUploadFileV1 interface{}

//GetListOfAvailableNamespacesV1 Get list of available namespaces - 3f89-bbfc-4f6b-8b50
/* Returns list of available namespaces



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-list-of-available-namespaces-v1
*/
func (s *FileService) GetListOfAvailableNamespacesV1() (*ResponseFileGetListOfAvailableNamespacesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/file/namespace"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseFileGetListOfAvailableNamespacesV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetListOfAvailableNamespacesV1()
		}
		return nil, response, fmt.Errorf("error with operation GetListOfAvailableNamespacesV1")
	}

	result := response.Result().(*ResponseFileGetListOfAvailableNamespacesV1)
	return result, response, err

}

//GetListOfFilesV1 Get list of files - 42b6-a86e-44b8-bdfc
/* Returns list of files under a specific namespace


@param nameSpace nameSpace path parameter. A listing of fileId's


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-list-of-files-v1
*/
func (s *FileService) GetListOfFilesV1(nameSpace string) (*ResponseFileGetListOfFilesV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/file/namespace/{nameSpace}"
	path = strings.Replace(path, "{nameSpace}", fmt.Sprintf("%v", nameSpace), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseFileGetListOfFilesV1{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetListOfFilesV1(nameSpace)
		}
		return nil, response, fmt.Errorf("error with operation GetListOfFilesV1")
	}

	result := response.Result().(*ResponseFileGetListOfFilesV1)
	return result, response, err

}

//DownloadAFileByFileIDV1 Download a file by fileId - 9698-c8ec-4a0b-8c1a
/* Downloads a file specified by fileId


@param fileID fileId path parameter. File Identification number


Documentation Link: https://developer.cisco.com/docs/dna-center/#!download-a-file-by-file-id-v1
*/
func (s *FileService) DownloadAFileByFileIDV1(fileID string) (FileDownload, *resty.Response, error) {
	path := "/dna/intent/api/v1/file/{fileId}"
	path = strings.Replace(path, "{fileId}", fmt.Sprintf("%v", fileID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		// SetResult(&ResponseFileDownloadAFileByFileIDV1{}).
		SetError(&Error).
		Get(path)

	fdownload := FileDownload{}
	if err != nil {
		return fdownload, nil, err
	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DownloadAFileByFileIDV1(fileID)
		}
		return fdownload, response, fmt.Errorf("error with operation ExportTrustedCertificate")
	}

	fdownload.FileData = response.Body()
	headerVal := response.Header()["Content-Disposition"][0]
	fname := strings.SplitAfter(headerVal, "=")
	fdownload.FileName = strings.ReplaceAll(fname[1], "\"", "")

	return fdownload, response, err

}

//UploadFileV1 uploadFile - 15bf-fb0f-44c8-98f2
/* Uploads a new file within a specific nameSpace


@param nameSpace nameSpace path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!upload-file-v1
*/
func (s *FileService) UploadFileV1(nameSpace string, UploadFileMultipartFields *UploadFileMultipartFields) (*ResponseFileUploadFileV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/file/{nameSpace}"
	path = strings.Replace(path, "{nameSpace}", fmt.Sprintf("%v", nameSpace), -1)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if UploadFileMultipartFields != nil {
		clientRequest = clientRequest.SetFileReader("file", UploadFileMultipartFields.FileName, UploadFileMultipartFields.File)
	}

	response, err = clientRequest.

		// SetResult(&ResponseFileUploadFileV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.UploadFileV1(nameSpace, UploadFileMultipartFields)
		}

		return nil, response, fmt.Errorf("error with operation UploadFileV1")
	}

	result := response.Result().(ResponseFileUploadFileV1)

	return &result, response, err

}

// Alias Function
/*
This method acts as an alias for the method `UploadFileV1`
*/
func (s *FileService) UploadFile(nameSpace string, UploadFileMultipartFields *UploadFileMultipartFields) (*ResponseFileUploadFileV1, *resty.Response, error) {
	return s.UploadFileV1(nameSpace, UploadFileMultipartFields)
}

// Alias Function
/*
This method acts as an alias for the method `GetListOfFilesV1`
*/
func (s *FileService) GetListOfFiles(nameSpace string) (*ResponseFileGetListOfFilesV1, *resty.Response, error) {
	return s.GetListOfFilesV1(nameSpace)
}

// Alias Function
/*
This method acts as an alias for the method `GetListOfAvailableNamespacesV1`
*/
func (s *FileService) GetListOfAvailableNamespaces() (*ResponseFileGetListOfAvailableNamespacesV1, *resty.Response, error) {
	return s.GetListOfAvailableNamespacesV1()
}

// Alias Function
/*
This method acts as an alias for the method `DownloadAFileByFileIDV1`
*/
func (s *FileService) DownloadAFileByFileID(fileID string) (FileDownload, *resty.Response, error) {
	return s.DownloadAFileByFileIDV1(fileID)
}
