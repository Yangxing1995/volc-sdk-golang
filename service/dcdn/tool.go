package dcdn

func (s *DCDN) SendCommonRequest(action string, in interface{}, out interface{}) error {
	return s.post(action, in, out)
}

func (s *DCDN) ValidateResponse(meta *ResponseMetadata) error {
	return validateResponse(meta)
}
