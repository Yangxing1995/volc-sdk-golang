package clb

func (s *CLB) SendCommonRequest(action string, in interface{}, out interface{}) error {
	return s.post(action, in, out)
}

func (s *CLB) ValidateResponse(meta *ResponseMetadata) error {
	return validateResponse(meta)
}
