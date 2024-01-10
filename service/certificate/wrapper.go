package certificate

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
