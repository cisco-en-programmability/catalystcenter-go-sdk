package catalyst

import (
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
)

type CiscoTrustedCertificatesService service

//ImportTrustedCertificateV1 Import Trusted Certificate - 65bd-0ac7-465a-8bbc
/* Imports trusted certificate into a truststore. Accepts .pem or .der file as input.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!import-trusted-certificate
*/
func (s *CiscoTrustedCertificatesService) ImportTrustedCertificateV1() (*resty.Response, error) {
	path := "/dna/intent/api/v1/trustedCertificates/import"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.ImportTrustedCertificateV1()
		}

		return response, fmt.Errorf("error with operation ImportTrustedCertificateV1")
	}

	return response, err

}

// Alias Function
/*
This method acts as an alias for the method `ImportTrustedCertificateV1`
*/
func (s *CiscoTrustedCertificatesService) ImportTrustedCertificate() (*resty.Response, error) {
	return s.ImportTrustedCertificateV1()
}
