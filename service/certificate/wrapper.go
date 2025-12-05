package certificate

// ImportCertificate 上传证书
// @https://www.volcengine.com/docs/6638/1365580?lang=zh
func (s *Certificate) ImportCertificate(dto *ImportCertificateRequest, options ...OptionArg) (responseBody *ImportCertificateResponse, err error) {
	responseBody = new(ImportCertificateResponse)
	if err = s.post("ImportCertificate", dto, responseBody, options...); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *Certificate) CertificateGetInstanceList(dto *CertificateGetInstanceListRequest, options ...OptionArg) (responseBody *CertificateGetInstanceListResponse, err error) {
	responseBody = new(CertificateGetInstanceListResponse)
	if err = s.post("CertificateGetInstanceList", dto, responseBody, options...); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *Certificate) CertificateGetInstance(dto *CertificateGetInstanceRequest, options ...OptionArg) (responseBody *CertificateGetInstanceResponse, err error) {
	responseBody = new(CertificateGetInstanceResponse)
	if err = s.post("CertificateGetInstance", dto, responseBody, options...); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}
