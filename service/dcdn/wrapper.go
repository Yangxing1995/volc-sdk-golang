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
