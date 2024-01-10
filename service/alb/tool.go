package alb

func (s *ALB) SendCommonRequest(action string, in interface{}, out interface{}) error {
	return s.post(action, in, out)
}

func (s *ALB) ValidateResponse(meta *ResponseMetadata) error {
	return validateResponse(meta)
}
