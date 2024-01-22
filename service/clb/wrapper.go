package clb

func (s *CLB) DescribeLoadBalancers(dto *DescribeLoadBalancersRequest, options ...OptionArg) (responseBody *DescribeLoadBalancersResponse, err error) {
	responseBody = new(DescribeLoadBalancersResponse)
	if err = s.post("DescribeLoadBalancers", dto, responseBody, options...); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CLB) DescribeLoadBalancerAttributes(dto *DescribeLoadBalancerAttributesRequest, options ...OptionArg) (responseBody *DescribeLoadBalancerAttributesResponse, err error) {
	responseBody = new(DescribeLoadBalancerAttributesResponse)
	if err = s.post("DescribeLoadBalancerAttributes", dto, responseBody, options...); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CLB) DescribeListeners(dto *DescribeListenersRequest, options ...OptionArg) (responseBody *DescribeListenersResponse, err error) {
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
func (s *CLB) ModifyListenerAttributes(dto *ModifyListenerAttributesRequest, options ...OptionArg) (responseBody *ModifyListenerAttributesResponse, err error) {
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
func (s *CLB) DescribeRules(dto *DescribeRulesRequest, options ...OptionArg) (responseBody *DescribeRulesResponse, err error) {
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
// 证书的格式要求 https://www.volcengine.com/docs/6406/68086
func (s *CLB) UploadCertificate(dto *UploadCertificateRequest, options ...OptionArg) (responseBody *UploadCertificateResponse, err error) {
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
func (s *CLB) DescribeCertificates(dto *DescribeCertificatesRequest, options ...OptionArg) (responseBody *DescribeCertificatesResponse, err error) {
	responseBody = new(DescribeCertificatesResponse)
	if err = s.post("DescribeCertificates", dto, responseBody, options...); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}
