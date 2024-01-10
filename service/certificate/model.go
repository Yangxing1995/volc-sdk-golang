package certificate

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

type BaseFindCond struct {
	PageNum  *int64 `json:",omitempty"`
	PageSize *int64 `json:",omitempty"`
}

type ImportCertificateRequest struct {
	Tag                 string           `json:"tag,omitempty"`
	Project             string           `json:"project,omitempty"`
	NoVerifyAndFixChain *bool            `json:"no_verify_and_fix_chain,omitempty"` // 跳过检查证书的合法性
	Repeatable          *bool            `json:"repeatable,omitempty"`
	CertificateInfo     *CertificateInfo `json:"certificate_info,omitempty"`

	GmCertificateInfo *GmCertificateInfo `json:"gm_certificate_info,omitempty"`
}

type CertificateInfo struct {
	Certificate string `json:"certificate"`
	PrivateKey  string `json:"private_key"`
}

type GmCertificateInfo struct {
	EncryptCertificate string `json:"encrypt_certificate"`
	EncryptPrivateKey  string `json:"encrypt_private_key"`
	SignCertificate    string `json:"sign_certificate"`
	SignPrivateKey     string `json:"sign_private_key"`
}

type ImportCertificateResponse struct {
	ResponseMetadata *ResponseMetadata        `json:",omitempty"`
	Result           *ImportCertificateResult `json:",omitempty"`
}

type ImportCertificateResult struct {
	ID       string `json:"id"`
	RepeatID string `json:"repeat_id"` // 重复证书的ID
}
