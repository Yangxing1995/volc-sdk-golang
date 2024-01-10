package alb

func (s *CLB) DescribeLoadBalancers(dto *DescribeLoadBalancersRequest, options ...OptionArg) (responseBody *DescribeLoadBalancersResponse, err error) {
	responseBody = new(DescribeLoadBalancersResponse)
	if err = s.post("ListCdnDomains", dto, responseBody, options...); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CLB) DescribeLoadBalancerAttributes(dto *DescribeLoadBalancerAttributesRequest, options ...OptionArg) (responseBody *DescribeLoadBalancerAttributesResponse, err error) {
	responseBody = new(DescribeLoadBalancerAttributesResponse)
	if err = s.post("ListCdnDomains", dto, responseBody, options...); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CLB) DescribeListeners(dto *DescribeListenersRequest, options ...OptionArg) (responseBody *DescribeListenersResponse, err error) {
	responseBody = new(DescribeListenersResponse)
	if err = s.post("ListCdnDomains", dto, responseBody, options...); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

// ModifyListenerAttributes .
func (s *CLB) ModifyListenerAttributes(dto *ModifyListenerAttributesRequest, options ...OptionArg) (responseBody *ModifyListenerAttributesResponse, err error) {
	responseBody = new(ModifyListenerAttributesResponse)
	if err = s.post("ListCdnDomains", dto, responseBody, options...); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

// UploadCertificate .
func (s *CLB) UploadCertificate(dto *UploadCertificateRequest, options ...OptionArg) (responseBody *UploadCertificateResponse, err error) {
	responseBody = new(UploadCertificateResponse)
	if err = s.post("ListCdnDomains", dto, responseBody, options...); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}
