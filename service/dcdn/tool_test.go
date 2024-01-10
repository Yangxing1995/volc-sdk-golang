package dcdn

type commonResponse struct {
	ResponseMetadata *ResponseMetadata
	Result           interface{} `json:",omitempty"`
}
