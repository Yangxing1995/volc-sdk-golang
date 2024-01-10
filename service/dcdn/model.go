package dcdn

type ErrorObj struct {
	CodeN   int64
	Code    string
	Message string
}

type ResponseMetadata struct {
	RequestId string
	Service   *string   `json:",omitempty"`
	Region    *string   `json:",omitempty"`
	Action    *string   `json:",omitempty"`
	Version   *string   `json:",omitempty"`
	Error     *ErrorObj `json:",omitempty"`
}

type DescribeUserDomainsRequest struct {
	PageNum  *int64    `json:",omitempty"`
	PageSize *int64    `json:",omitempty"`
	Project  *[]string `json:",omitempty"`
}

type DescribeUserDomainsResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeUserDomainsResult
}

type DescribeUserDomainsResult struct {
	AllDomainNum    int `json:"AllDomainNum"`
	OnlineDomainNum int `json:"OnlineDomainNum"`
	Domains         []struct {
		Domain string `json:"Domain"`
		Status string `json:"Status"`
		Scope  string `json:"Scope"`
	} `json:"Domains"`
	PageNum  int `json:"PageNum"`
	PageSize int `json:"PageSize"`
}

type CreateCertBindRequest struct {
	CertID    string   `json:"CertId"`
	DomainIds []string `json:"DomainIds"`
}

type CreateCertBindResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
}
