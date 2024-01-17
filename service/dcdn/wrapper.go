package dcdn

func (s *DCDN) DescribeUserDomains(dto *DescribeUserDomainsRequest, options ...OptionArg) (responseBody *DescribeUserDomainsResponse, err error) {
	responseBody = new(DescribeUserDomainsResponse)
	if err = s.post("DescribeUserDomains", dto, responseBody, options...); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *DCDN) DescribeDomainConfig(dto *DescribeDomainConfigRequest, options ...OptionArg) (responseBody *DescribeDomainConfigResponse, err error) {
	responseBody = new(DescribeDomainConfigResponse)
	if err = s.post("DescribeDomainDetail", dto, responseBody, options...); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *DCDN) DescribeDomainDetail(dto *DescribeDomainDetailRequest, options ...OptionArg) (responseBody *DescribeDomainDetailResponse, err error) {
	responseBody = new(DescribeDomainDetailResponse)
	if err = s.post("DescribeDomainDetail", dto, responseBody, options...); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

// UploadSelfCert 还未开放
func (s *DCDN) UploadSelfCert(dto *UploadSelfCertRequest, options ...OptionArg) (responseBody *UploadSelfCertResponse, err error) {
	responseBody = new(UploadSelfCertResponse)
	if err = s.post("UploadSelfCert", dto, responseBody, options...); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

// CreateCertBind 此API可能存在问题, 请谨慎使用
func (s *DCDN) CreateCertBind(dto *CreateCertBindRequest, options ...OptionArg) (responseBody *CreateCertBindResponse, err error) {
	responseBody = new(CreateCertBindResponse)
	if err = s.post("CreateCertBind", dto, responseBody, options...); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}
