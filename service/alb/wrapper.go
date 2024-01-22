package alb

func (s *ALB) DescribeLoadBalancers(dto *DescribeLoadBalancersRequest, options ...OptionArg) (responseBody *DescribeLoadBalancersResponse, err error) {
	responseBody = new(DescribeLoadBalancersResponse)
	if err = s.post("DescribeLoadBalancers", dto, responseBody, options...); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *ALB) DescribeLoadBalancerAttributes(dto *DescribeLoadBalancerAttributesRequest, options ...OptionArg) (responseBody *DescribeLoadBalancerAttributesResponse, err error) {
	responseBody = new(DescribeLoadBalancerAttributesResponse)
	if err = s.post("DescribeLoadBalancerAttributes", dto, responseBody, options...); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *ALB) DescribeListeners(dto *DescribeListenersRequest, options ...OptionArg) (responseBody *DescribeListenersResponse, err error) {
	responseBody = new(DescribeListenersResponse)
	if err = s.post("DescribeListeners", dto, responseBody, options...); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

// ModifyListenerAttributes .
func (s *ALB) ModifyListenerAttributes(dto *ModifyListenerAttributesRequest, options ...OptionArg) (responseBody *ModifyListenerAttributesResponse, err error) {
	responseBody = new(ModifyListenerAttributesResponse)
	if err = s.post("ModifyListenerAttributes", dto, responseBody, options...); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

// DescribeRules .
func (s *ALB) DescribeRules(dto *DescribeRulesRequest, options ...OptionArg) (responseBody *DescribeRulesResponse, err error) {
	responseBody = new(DescribeRulesResponse)
	if err = s.post("DescribeRules", dto, responseBody, options...); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

// UploadCertificate .
func (s *ALB) UploadCertificate(dto *UploadCertificateRequest, options ...OptionArg) (responseBody *UploadCertificateResponse, err error) {
	responseBody = new(UploadCertificateResponse)
	if err = s.post("UploadCertificate", dto, responseBody, options...); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

// DescribeCertificates .
func (s *ALB) DescribeCertificates(dto *DescribeCertificatesRequest, options ...OptionArg) (responseBody *DescribeCertificatesResponse, err error) {
	responseBody = new(DescribeCertificatesResponse)
	if err = s.post("DescribeCertificates", dto, responseBody, options...); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}
