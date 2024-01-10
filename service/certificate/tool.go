package certificate

func (s *Certificate) SendCommonRequest(action string, in interface{}, out interface{}) error {
	return s.post(action, in, out)
}

func (s *Certificate) ValidateResponse(meta *ResponseMetadata) error {
	return validateResponse(meta)
}
